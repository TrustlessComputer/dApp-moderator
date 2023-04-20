import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "hardhat-deploy";
import "@ericxstone/hardhat-blockscout-verify";


const config: HardhatUserConfig = {
    defaultNetwork: "mynw",
    solidity: {
      compilers: [
        { version: "0.5.16", settings: {} },
        { version: "0.8.17", settings: {} },
      ]
    },
    networks: {
      mynw: {
        url: "http://127.0.0.1:8545",
        accounts: {
          mnemonic: "",
        },
        // issue: https://github.com/NomicFoundation/hardhat/issues/3136
        // workaround: https://github.com/NomicFoundation/hardhat/issues/2672#issuecomment-1167409582
        timeout: 100_000,
      },
    },
  };
  
  export default config;