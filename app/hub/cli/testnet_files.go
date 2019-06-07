package cli

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	serverConfig "github.com/cosmos/cosmos-sdk/server/config"
	csdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tmConfig "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/types"
	tmTime "github.com/tendermint/tendermint/types/time"

	app "github.com/ironman0x7b2/sentinel-sdk/app/hub"
)

const nodeDirPerm = 0755

// nolint:gochecknoglobals
var (
	flagNodeDirPrefix     = "node-dir-prefix"
	flagNumValidators     = "v"
	flagOutputDir         = "output-dir"
	flagNodeDaemonHome    = "node-daemon-home"
	flagNodeCliHome       = "node-cli-home"
	flagStartingIPAddress = "starting-ip-address"
)

func TestNetFilesCmd(ctx *server.Context, cdc *codec.Codec) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "testnet",
		Short: "Initialize files for a hubd testnet",
		Long: `testnet will create "v" number of directories and populate each with
necessary files (private validator, genesis, config, etc.).

Note, strict routability for addresses is turned off in the config file.

Example:
	hubd testnet --v 4 --output-dir ./output --starting-ip-address 192.168.10.2
	`,
		RunE: func(_ *cobra.Command, _ []string) error {
			config := ctx.Config
			return initTestNet(config, cdc)
		},
	}

	cmd.Flags().Int(flagNumValidators, 4,
		"Number of validators to initialize the testnet with",
	)
	cmd.Flags().StringP(flagOutputDir, "o", "./mytestnet",
		"Directory to store initialization data for the testnet",
	)
	cmd.Flags().String(flagNodeDirPrefix, "node",
		"Prefix the directory name for each node with (node results in node0, node1, ...)",
	)
	cmd.Flags().String(flagNodeDaemonHome, "hubd",
		"Home directory of the node's daemon configuration",
	)
	cmd.Flags().String(flagNodeCliHome, "hubcli",
		"Home directory of the node's cli configuration",
	)
	cmd.Flags().String(flagStartingIPAddress, "192.168.0.1",
		"Starting IP address (192.168.0.1 results in persistent peers list ID0@192.168.0.1:46656, ID1@192.168.0.2:46656, ...)") // nolint:lll

	cmd.Flags().String(
		client.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created",
	)
	cmd.Flags().String(
		server.FlagMinGasPrices, fmt.Sprintf("0.000006%s", "stake"),
		"Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.001stake)",
	)

	return cmd
}

// nolint:gocyclo
func initTestNet(config *tmConfig.Config, cdc *codec.Codec) error {
	var chainID string

	outDir := viper.GetString(flagOutputDir)
	numValidators := viper.GetInt(flagNumValidators)

	chainID = viper.GetString(client.FlagChainID)
	if chainID == "" {
		chainID = "chain-" + common.RandStr(6)
	}

	monikers := make([]string, numValidators)
	nodeIDs := make([]string, numValidators)
	valPubKeys := make([]crypto.PubKey, numValidators)

	hubConfig := serverConfig.DefaultConfig()
	hubConfig.MinGasPrices = viper.GetString(server.FlagMinGasPrices)

	accounts := make([]app.GenesisAccount, 0, numValidators)
	genFiles := make([]string, 0, numValidators)

	for i := 0; i < numValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", viper.GetString(flagNodeDirPrefix), i)
		nodeDaemonHomeName := viper.GetString(flagNodeDaemonHome)
		nodeCliHomeName := viper.GetString(flagNodeCliHome)
		nodeDir := filepath.Join(outDir, nodeDirName, nodeDaemonHomeName)
		clientDir := filepath.Join(outDir, nodeDirName, nodeCliHomeName)
		genTxsDir := filepath.Join(outDir, "gentxs")

		config.SetRoot(nodeDir)

		err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm)
		if err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}

		err = os.MkdirAll(clientDir, nodeDirPerm)
		if err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}

		monikers = append(monikers, nodeDirName)
		config.Moniker = nodeDirName

		ip, err := getIP(i, viper.GetString(flagStartingIPAddress))
		if err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}

		nodeIDs[i], valPubKeys[i], err = InitializeNodeValidatorFiles(config)
		if err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}

		memo := fmt.Sprintf("%s@%s:26656", nodeIDs[i], ip)
		genFiles = append(genFiles, config.GenesisFile())

		buf := client.BufferStdin()
		prompt := fmt.Sprintf(
			"Password for account '%s' (default %s):", nodeDirName, app.DefaultKeyPass,
		)

		keyPass, err := client.GetPassword(prompt, buf)
		if err != nil && keyPass != "" {
			return err
		}

		if keyPass == "" {
			keyPass = app.DefaultKeyPass
		}

		addr, secret, err := server.GenerateSaveCoinKey(clientDir, nodeDirName, keyPass, true)
		if err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}

		info := map[string]string{"secret": secret}

		cliPrint, err := json.Marshal(info)
		if err != nil {
			return err
		}

		err = writeFile(fmt.Sprintf("%v.json", "key_seed"), clientDir, cliPrint)
		if err != nil {
			return err
		}

		accTokens := csdk.TokensFromTendermintPower(1000)
		accStakingTokens := csdk.TokensFromTendermintPower(500)
		accounts = append(accounts, app.GenesisAccount{
			Address: addr,
			Coins: csdk.Coins{
				csdk.NewCoin(fmt.Sprintf("%stoken", nodeDirName), accTokens),
				csdk.NewCoin("stake", accStakingTokens),
			},
		})

		valTokens := csdk.TokensFromTendermintPower(100)
		msg := staking.NewMsgCreateValidator(
			csdk.ValAddress(addr),
			valPubKeys[i],
			csdk.NewCoin("stake", valTokens),
			staking.NewDescription(nodeDirName, "", "", ""),
			staking.NewCommissionMsg(csdk.ZeroDec(), csdk.ZeroDec(), csdk.ZeroDec()),
			csdk.OneInt(),
		)

		kb, err := keys.NewKeyBaseFromDir(clientDir)
		if err != nil {
			return err
		}

		tx := auth.NewStdTx([]csdk.Msg{msg}, auth.StdFee{}, []auth.StdSignature{}, memo)
		txBuilder := authTxBuilder.NewTxBuilderFromCLI().WithChainID(chainID).WithMemo(memo).WithKeybase(kb)

		signedTx, err := txBuilder.SignStdTx(nodeDirName, app.DefaultKeyPass, tx, false)
		if err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}

		txBytes, err := cdc.MarshalJSON(signedTx)
		if err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}

		err = writeFile(fmt.Sprintf("%v.json", nodeDirName), genTxsDir, txBytes)
		if err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}

		hubConfigFilePath := filepath.Join(nodeDir, "config/hubd.toml")
		serverConfig.WriteConfigFile(hubConfigFilePath, hubConfig)
	}

	if err := initGenFiles(cdc, chainID, accounts, genFiles, numValidators); err != nil {
		return err
	}

	err := collectGenFiles(
		cdc, config, chainID, monikers, nodeIDs, valPubKeys, numValidators,
		outDir, viper.GetString(flagNodeDirPrefix), viper.GetString(flagNodeDaemonHome),
	)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully initialized %d node directories\n", numValidators)
	return nil
}

func initGenFiles(cdc *codec.Codec, chainID string, accs []app.GenesisAccount,
	genFiles []string, numValidators int) error {

	appGenState := app.NewDefaultGenesisState()
	appGenState.Accounts = accs

	appGenStateJSON, err := codec.MarshalJSONIndent(cdc, appGenState)
	if err != nil {
		return err
	}

	genDoc := types.GenesisDoc{
		ChainID:    chainID,
		AppState:   appGenStateJSON,
		Validators: nil,
	}

	for i := 0; i < numValidators; i++ {
		if err := genDoc.SaveAs(genFiles[i]); err != nil {
			return err
		}
	}

	return nil
}

func collectGenFiles(cdc *codec.Codec, config *tmConfig.Config, chainID string,
	monikers, nodeIDs []string, valPubKeys []crypto.PubKey,
	numValidators int, outDir, nodeDirPrefix, nodeDaemonHomeName string) error {

	var appState json.RawMessage
	genTime := tmTime.Now()

	for i := 0; i < numValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", nodeDirPrefix, i)
		nodeDir := filepath.Join(outDir, nodeDirName, nodeDaemonHomeName)
		genTxsDir := filepath.Join(outDir, "gentxs")
		moniker := monikers[i]
		config.Moniker = nodeDirName

		config.SetRoot(nodeDir)

		nodeID, valPubKey := nodeIDs[i], valPubKeys[i]
		initCfg := newInitConfig(chainID, genTxsDir, moniker, nodeID, valPubKey)

		genDoc, err := loadGenesisDoc(cdc, config.GenesisFile())
		if err != nil {
			return err
		}

		nodeAppState, err := genAppStateFromConfig(cdc, config, initCfg, genDoc)
		if err != nil {
			return err
		}

		if appState == nil {
			appState = nodeAppState
		}

		genFile := config.GenesisFile()

		err = ExportGenesisFileWithTime(genFile, chainID, nil, appState, genTime)
		if err != nil {
			return err
		}
	}

	return nil
}

func getIP(i int, startingIPAddr string) (string, error) {
	var (
		ip  string
		err error
	)

	if len(startingIPAddr) == 0 {
		ip, err = server.ExternalIP()
		if err != nil {
			return "", err
		}
	} else {
		ip, err = calculateIP(startingIPAddr, i)
		if err != nil {
			return "", err
		}
	}

	return ip, nil
}

func writeFile(name string, dir string, contents []byte) error {
	writePath := filepath.Join(dir)
	file := filepath.Join(writePath, name)

	err := common.EnsureDir(writePath, 0700)
	if err != nil {
		return err
	}

	err = common.WriteFile(file, contents, 0600)
	if err != nil {
		return err
	}

	return nil
}

func calculateIP(ip string, i int) (string, error) {
	ipv4 := net.ParseIP(ip).To4()
	if ipv4 == nil {
		return "", fmt.Errorf("%v: non ipv4 address", ip)
	}

	for j := 0; j < i; j++ {
		ipv4[3]++
	}

	return ipv4.String(), nil
}
