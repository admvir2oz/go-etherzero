// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

var MainnetBootnodes = []string{
	"enode://f6fc7cc448a2628cd635612118a2de0dd7a41a2f343fc554900d43f38d1b5bee71a46ac5b8fb743fd054477866bf6a9f8de3e4609e2be0fed8c813ba87476d6b@206.189.151.92:20001", //
	"enode://34f90447b8f4111827eefbe484fd7957123422745535ac9ef91f9f24862151c42b062b19292ff26878982bd527ab5636b29b8b472c1e34a9bda6921b8416d35f@206.189.151.92:20002", //
	"enode://c2093b60452ef3b28217365ed090e4cc6e42de5403f1937133f9a5654bd0ba019c76ac2383c0365a67f32e8cf3052ce824a3e2fa69879ab653751105a670e3b5@142.93.57.240:20003", //
	"enode://d6f33d7be72efc503fa92ad3c82ec250fd42010dabaa02760e356c1d27384e77b70086c4a0ede0d22f3c8c03c26c5c17741b396a2beba19714dca273bcbb8474@142.93.57.240:20004", //
	"enode://3d2b9d5215a2c542339b2dba860b7c962943fde7c960a43414911fcf9d4907da2e53f11bf950d9d74a87c01ce9498b92b96a4b0a9d10b4a1f9c9774c14ce07e8@206.189.105.82:20005", //
	"enode://5fbf241a0fad4f4452e6ac9e1d1343db707bc09a86f8c3d5ff2ee84c6f33f9843331051b7788d39e17df4ac65d31946c2461215d2bedd13d0bdaeb2147fcb7aa@206.189.105.82:20006", //
	"enode://0b7016dd9abad235c009c127b5b2b0d05d2078c579b27e2529212c4bcf34914c9a4fa4ce3273ca05988ea1330bc1b85adde285a9e89ec763b521ed93efb52a78@209.97.171.164:20007", //
	"enode://51193b58eebf8b0614a2d05d223d376a8492a329dc1f9d475be2902009887a82e9faa0d059eba2c75c1967488a215e5baf4aa6cea2df22455fbcde90ec9dc6d6@209.97.171.164:20008", //
	"enode://2101572dabc573348e3f31a83bcc6bab8531481f8904556716754ca4a57e93cce10f5b2a3ab1f54d21c763e262a88dcef355d03e9de8152f472d2ab1ec0ea284@209.97.172.45:20009", //
	"enode://c15e1dddbda67656293d932ee16efef44e11fd3a57af428ce63ec24e92e09dc04b0448d35a56cdddec7520e4109d8d0e53a5d00ec04758506e133f65d8f52d3c@209.97.172.45:20010", //
	"enode://6c667f2ce11ccde49e868af71cf1ec20149a1edc3e67722896057bbf6356d5f3352b10feb0cb8f3e97ece3c8a15d0493df62bdc87f977a98c45de5f8bef0c0e7@142.93.45.4:20011", //
	"enode://7d86aa13716729476692123fcc6f2a23a1abc64eaef8cca87d823465b64ba013d2350c2439c1c91a3bb1bb8ae856f8e15bff8b0ff0378f96c42ade7a710aa9ad@142.93.45.4:20012", //
	"enode://99a488ec68a7c57de542f930c6537f96ba1c3db8ac3f6774bb62ae4ba77519a7c4b93a7c80ee99e5d9ac7b329fc3f0ebf80c49e7616ecde29b5e42a9c147bc72@178.128.195.117:20013", //
	"enode://14af195555e043442859e768037c761d856067079b0cece9a6e1c0335b6724762776c880ac9f76aa16e782c92797315b0c7394ffae490e208386a4c311f15f76@178.128.195.117:20014", //
	"enode://79585df29910656ace1ad113fc4668e01bbf9fbe0ffaeba126771013dd388169fc7c880ab287a35ae9301bf185abfbd98f3a27757e5f7b64250dc017230be7bf@165.227.40.25:20015", //
	"enode://7eeb03ba746f2177f71ce6e27331156e4ae04cef7143c01e7daa014e61b20e1b7822b1056af418242ee7805f11c9862282ded65c85d1d87764517458a2070b87@165.227.40.25:20016", //
	"enode://129a841e7c2fef28d5e086fc750b69f538e4c00f1742b0b71dbc7a837dccb39b97624b9b17e602c044940ae84f2d8dd25a7983fde9083ec230dae5e044e5e5f0@139.59.19.25:20017", //
	"enode://0125583fa22682d37723b55835db06c7eb934215bff5b8809cda6aa31f6fce48f85bbaf9c7cbc33f1440309a86b2710c758368aa9ccef47a53dc4c0324e25077@139.59.19.25:20018", //
	"enode://c325030337695371f3103764e006b164d64839addc3b101bfc936f97e77dd6076022d43b442e75e3ff21d209d6804cf95e54dbdbb9536de5c11704c948358443@178.128.121.124:20019", //
	"enode://58291c8c1cddc7f4394f8b2882ce2dc73d77412088d22366389186e006e3df5314e3cc3dc6a1072d221f84ead2ccd2d62a266affff7df5bf4cfc25f4410b4785@178.128.121.124:20020", //
}

var MainnetMasternodes = []string{
	"enode://f6fc7cc448a2628cd635612118a2de0dd7a41a2f343fc554900d43f38d1b5bee71a46ac5b8fb743fd054477866bf6a9f8de3e4609e2be0fed8c813ba87476d6b", //
	"enode://34f90447b8f4111827eefbe484fd7957123422745535ac9ef91f9f24862151c42b062b19292ff26878982bd527ab5636b29b8b472c1e34a9bda6921b8416d35f", //
	"enode://c2093b60452ef3b28217365ed090e4cc6e42de5403f1937133f9a5654bd0ba019c76ac2383c0365a67f32e8cf3052ce824a3e2fa69879ab653751105a670e3b5", //
	"enode://d6f33d7be72efc503fa92ad3c82ec250fd42010dabaa02760e356c1d27384e77b70086c4a0ede0d22f3c8c03c26c5c17741b396a2beba19714dca273bcbb8474", //
	"enode://3d2b9d5215a2c542339b2dba860b7c962943fde7c960a43414911fcf9d4907da2e53f11bf950d9d74a87c01ce9498b92b96a4b0a9d10b4a1f9c9774c14ce07e8", //
	"enode://5fbf241a0fad4f4452e6ac9e1d1343db707bc09a86f8c3d5ff2ee84c6f33f9843331051b7788d39e17df4ac65d31946c2461215d2bedd13d0bdaeb2147fcb7aa", //
	"enode://0b7016dd9abad235c009c127b5b2b0d05d2078c579b27e2529212c4bcf34914c9a4fa4ce3273ca05988ea1330bc1b85adde285a9e89ec763b521ed93efb52a78", //
	"enode://51193b58eebf8b0614a2d05d223d376a8492a329dc1f9d475be2902009887a82e9faa0d059eba2c75c1967488a215e5baf4aa6cea2df22455fbcde90ec9dc6d6", //
	"enode://2101572dabc573348e3f31a83bcc6bab8531481f8904556716754ca4a57e93cce10f5b2a3ab1f54d21c763e262a88dcef355d03e9de8152f472d2ab1ec0ea284", //
	"enode://c15e1dddbda67656293d932ee16efef44e11fd3a57af428ce63ec24e92e09dc04b0448d35a56cdddec7520e4109d8d0e53a5d00ec04758506e133f65d8f52d3c", //
	"enode://6c667f2ce11ccde49e868af71cf1ec20149a1edc3e67722896057bbf6356d5f3352b10feb0cb8f3e97ece3c8a15d0493df62bdc87f977a98c45de5f8bef0c0e7", //
	"enode://7d86aa13716729476692123fcc6f2a23a1abc64eaef8cca87d823465b64ba013d2350c2439c1c91a3bb1bb8ae856f8e15bff8b0ff0378f96c42ade7a710aa9ad", //
	"enode://99a488ec68a7c57de542f930c6537f96ba1c3db8ac3f6774bb62ae4ba77519a7c4b93a7c80ee99e5d9ac7b329fc3f0ebf80c49e7616ecde29b5e42a9c147bc72", //
	"enode://14af195555e043442859e768037c761d856067079b0cece9a6e1c0335b6724762776c880ac9f76aa16e782c92797315b0c7394ffae490e208386a4c311f15f76", //
	"enode://79585df29910656ace1ad113fc4668e01bbf9fbe0ffaeba126771013dd388169fc7c880ab287a35ae9301bf185abfbd98f3a27757e5f7b64250dc017230be7bf", //
	"enode://7eeb03ba746f2177f71ce6e27331156e4ae04cef7143c01e7daa014e61b20e1b7822b1056af418242ee7805f11c9862282ded65c85d1d87764517458a2070b87", //
	"enode://129a841e7c2fef28d5e086fc750b69f538e4c00f1742b0b71dbc7a837dccb39b97624b9b17e602c044940ae84f2d8dd25a7983fde9083ec230dae5e044e5e5f0", //
	"enode://0125583fa22682d37723b55835db06c7eb934215bff5b8809cda6aa31f6fce48f85bbaf9c7cbc33f1440309a86b2710c758368aa9ccef47a53dc4c0324e25077", //
	"enode://c325030337695371f3103764e006b164d64839addc3b101bfc936f97e77dd6076022d43b442e75e3ff21d209d6804cf95e54dbdbb9536de5c11704c948358443", //
	"enode://58291c8c1cddc7f4394f8b2882ce2dc73d77412088d22366389186e006e3df5314e3cc3dc6a1072d221f84ead2ccd2d62a266affff7df5bf4cfc25f4410b4785", //

}
var TestnetMasternodes = []string{
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:21212",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:21212",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:21212", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:21212", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://c84c9860a017cadda359c2b63c29555811d02bf5839938107878ccce856447f67cc72adbad6837c18f823f4ee0a29d48405082ffb47fd490fa5d8d9f80b8ae78@127.0.0.1:21212", //
	"enode://7428587a4839162af1c2bc93629bcf45162abae8772ad93027dfdd94fddc7c653e085b5f58fd0def2533ca61a3e380f4f5e9e891d26eea8cc11df7dcf4b77189@127.0.0.1:21213", //
	"enode://fc7f473e7ccbdee93fce0749c4df890a8c3dd81a79678a939fb622f061ee48112d3cde8ed41f4000a21a21084ca382e339b5bd46eb58afa7962ee0f638d3342c@127.0.0.1:21214", //
	"enode://29416b59c5ce177413ab98f03fb307928e477c5437ec658cdd127633b19dc3968399990a58557e62d6ee843bb93f0d18639e42f3ec7ba0ec0869507d8b6ea551@127.0.0.1:21215", //
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:21212",
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",
}
