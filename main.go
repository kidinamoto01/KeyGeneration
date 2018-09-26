package main

import (
	"encoding/json"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"os"
	"io/ioutil"
	"fmt"
	"path"
)

// genesis piece structure for creating combined genesis
type GenesisTx struct {
	NodeID    string                   `json:"node_id"`
	IP        string                   `json:"ip"`
	Validator tmtypes.GenesisValidator `json:"validator"`
	AppGenTx  json.RawMessage          `json:"app_gen_tx"`
}


func AddCodec() *wire.Codec {
	cdc := amino.NewCodec()
	//cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	//cdc.RegisterConcrete(GenesisTx{}, "test/genesis", nil)
	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(ed25519.PubKeyEd25519{},
		"tendermint/PubKeyEd25519", nil)
	cdc.RegisterConcrete(secp256k1.PubKeySecp256k1{},
		"tendermint/PubKeySecp256k1", nil)
	return cdc
}


func main(){

	path := "/Users/suyu/go/src/github.com/irisnet/testnets/fuxi/fuxi-3000/config/gentx"

	cdc := AddCodec()

	ProcessGenTxs(path,cdc)
}


// append a genesis-piece
func ProcessGenTxs(genTxsDir string, cdc *wire.Codec) (err error) {

	var fos []os.FileInfo
	fos, err = ioutil.ReadDir(genTxsDir)
	if err != nil {
		return
	}

	//genTxs := make(map[string]crypto.Address)
	genTxs := map[string]crypto.Address{}
	//var nodeIDs []string
	fmt.Println(len(fos))
	for _, fo := range fos {
		filename := path.Join(genTxsDir, fo.Name())
		fmt.Println(fo.Name())
		if !fo.IsDir() && (path.Ext(filename) != ".json") {
			continue
		}

		// get the genTx
		var bz []byte
		bz, err = ioutil.ReadFile(filename)

		if err != nil {
			return
		}
		var genTx GenesisTx
		err = cdc.UnmarshalJSON(bz, &genTx)

		if err != nil {
			return
		}
		//		var val  tmtypes.GenesisValidator
		genTxs[fo.Name()] = genTx.Validator.PubKey.Address()

		//fmt.Println("got",genTx.Validator.PubKey.Address())

	}

	err = SaveAs(genTxs,"/Users/suyu/Documents/bianjie/fuxi/mystruct.json",cdc)
	if err!= nil{
		println(err)
	}

	return
}


// SaveAs is a utility method for saving GenensisDoc as a JSON file.
func SaveAs(genDoc interface{},filepath string, cdc *wire.Codec,) error {
	genDocBytes, err := cdc.MarshalJSONIndent(genDoc, "", "  ")

	//appGenTx := json.RawMessage(genDocBytes)

	if err != nil {
		return err
	}
	//cmn.WriteFile(file, genDocBytes, 0644)

	return ioutil.WriteFile(filepath, genDocBytes, 0644)
}



