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

import "github.com/dominant-strategies/go-quai/common"

// ColosseumBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the Colosseum test network.
var ColosseumBootnodes = map[string][]string{
	"prime": {
		"enode://3ccf9f8c6737ab6e73ff628695c5c00d5d46fa2dfe7a65b689e15c3027753e7343ecfd80d58022d8646972a67ac84305913b61aba6dbcb27a019488095c2a9b9@34.139.206.86",
		"enode://0b5b3ce249b688432a40cc514f6adde70acb412269c0d9a1f51877e179bd069ea4e6f1b984c3bb4bebca7eca576e022ec075a52da6291361771fd669546e9fca@34.168.1.153",
		"enode://e34d26c4de167d4f19de7d612697ec18ebf73f4c84767ecb0fed7bf9e0baf68a875033c085a7c9e13c75c404ca1f14a6260d4790552a9f6af6a1b0b50ed256af@34.168.35.191",
		"enode://63d05bbd8546c051c8df9e05b24251b3685b806da4ba314737918fe39f690fcd0ca074f52120ba3ed2edf90be5e4eed1e2505d6ab3041fef9d3a8e7d117e521f@34.82.221.18",
		"enode://d7af88081495f8a8e444de81cee6467e8695a6ed26de456382de40bec5cc3273d0b8d7d8b21091aac3dea6442b7368e9fd84717aed0fba9387a1af4bafd34bf6@35.184.84.143",
		"enode://840b36e3ccda764065992217157531b4c99be1b6a82d30567f7accf61600bc1253b0af60cab4d1d2e45b0b7d1142e2f23e5d9c817c7b6aa69357727739d7e218@35.196.20.106",
		"enode://80c67a2b859d4d7c23aada8f26ca17dba006efe4bc85868e58f3af4be4fbdad9fb96fe2b52a84e305d69588463d9f51ce2862b8300964242d75caa9f53bb8f82@35.202.3.128",
		"enode://512fad72315c674e935bb46a80068b2994134d196c9c98910775b2026d9bd04a2e0d63e2e5c998705c01396bdbfc32ba40e660ee938b65d49d92e38c03f2a586@35.224.101.46",
		"enode://c4d379b918023bdd45d57801597b55cce9992442dbddc1ed42c2ed4e4672e8d5faba47cfab80aec71a245fa5be3323e48fded89dbab422e4c01b22d3951b45d7@35.231.200.19",
	},
	"cyprus": {
		"enode://eaca846cacf4daa221abd910d53b3d0aec7cb781329d0d9a33002437c07f423bebed571c145f27dcb54ea821d6c2ad55c6778f408bc7a0ca33db6aab0093ebdd@34.139.206.86",
		"enode://adfe25ee467ce0ea9f8d8d33fce70d0c19593ffdde67b004e56b53f08fc39273b891c964d6a59bb514c44ead3591fec42e9b8b65e9c46833a6f2c3253fb0b99c@34.168.1.153",
		"enode://b896da42bf0ac68ec6ef98e9dfbe872603ee81a329042fc3d5a2d3996826af78bc0f7483caaa6ac298036f2a924c7b3539f29df3e63f9b101b4dc5779048a32b@34.168.35.191",
		"enode://29503623026075c643daa3d2e5b83bdb08dc7b2de52e8a201bcdcd8efd86bd61db43030b34eeb0bbc651a8217e51cdde204a7f8da9540e40ede508c7a8c82d04@34.82.221.18",
		"enode://b98e9f51e0afcba2d42d35c708cb69bbf15f2aae0160a49e7c32a43098c3df3a974e294d25cf7bc6be02b25257a074fa555daad1a2982007ac7ffad8b6b09ed6@35.184.84.143",
		"enode://6b67a6fdd89fe33ff9cc66c64360f7601789ee2badd1e37d4098d33526d97e6f5bce3e4f655ec99f36b37f9d8e1d17c611970530361b6ca0c2caf50f143be42d@35.196.20.106",
		"enode://860a7a6563bbfec6a8c0c3960a482c757a638442d1ffdf4711fd1b70a899fa3937db013363fbf21d0b0a8c8fa557da3509250a5391e4f3ff2efaf780f2b41a6c@35.202.3.128",
		"enode://8b490a284014ea09caf36d0c3b349282708991fab981c832c47977bf45898dbf279942e0b2298efbacb7a6e99725e63e93c24611f82afbcbce533cea5abdd57d@35.224.101.46",
		"enode://5774ae8d95803161e986ca5e3714c1d155e9bacd64e86e0b610a14c033a8bda870cad410e80b520d08e23a4c07aed78d04dfc53b101007aacc31e691f4dfc53c@35.231.200.19",
	},
	"paxos": {
		"enode://089a3ecef24fa28e1a8c914195fb9d8c68fe0ad9efa8453baa437d2c978720719995d669746c16d451be03e670a5e071b4235788450c78bf73802970ceabb46c@34.139.206.86",
		"enode://14f4f4f74c450c2ba2bd1bb6dba786695a592f00f646c25fc164e3c7c3e803dc4a364d241d89403bfb9a5c50924c5c59bcce742ea6d5aa6f3153ce2c09fd4e93@34.168.1.153",
		"enode://2862ade6942027dad0f5e4f2e184dfd5d9cc36d3c3b02905bd7e48d5665ac31365fa019d1933f331f629c7ad118c796dd3fffd6ec1e252c944e9184caa260d9b@34.168.35.191",
		"enode://a7e26ab06f055d9c1269d0740f2213da98108b81e91cda2d210d8465cdd118034897923d2d1a7b36be470e32b2fffd706bce5f78ea61fcc51c8a2a746ac11d46@34.82.221.18",
		"enode://efedf681345fb413c883d027182180da7453b5fff064c24481ad121707d30f2adcfb7676dfef10e306cd46e9869b958d4b681e81b00884e24080078546ce701d@35.184.84.143",
		"enode://b788ef076afe1e5632d9baee154eca3c0a0598756e26a9b4e7b90d4650c9a6a2f2d6af100e1ff2f444051f1b31143acbad9046fb84e5f2d00640ac0cc37126b4@35.196.20.106",
		"enode://0dc54a1bd9719157f07b963c4504e84c13ae0a7315b245067e75b80a6a5851b7c5da086ca4459ab5a236e621193dd9ef86e16faf616ba4147c3bdcd48a64f59b@35.202.3.128",
		"enode://118c6baa89b4e6fa13662add2fdb5674289dc68a58ce65ae36dfa69762746c29b44135c51b78cccf93a05f4124812a155bd8fd4bbb09b6d687f113d4fe47ccb8@35.224.101.46",
		"enode://fd11ffb765b80acafc20c7a134d1da52268f172de16bde25e3d7aa58918770956eadad72cd782cca816b33fee25ff8edf60631cf59f945f46e048ee3a7cd1f27@35.231.200.19",
	},
	"hydra": {
		"enode://801ca16bf9b0544e97033e9b7ca64936bfbc7bf255ce4362bcec92e68929e0195ed4bff44d436e5622cefa5f2ba25ac04de360124f04102337bad51e45e5dd67@34.139.206.86",
		"enode://88b409183fe652318026798db38039d4b2201b055700443a077c33531aafb646013c0590c58893a6ca1bc2571da8e492e9a386958f298765116e592a367eb2bd@34.168.1.153",
		"enode://44eef71ec4b10f8f809dae57d34520d253b535f9d58419aa98ff35d37a6d88c2037e98539f93fa8d580266e04d15dd98b34e9d301f9efaa827a77536a3792145@34.168.35.191",
		"enode://766608028cde8d4e8e4d84c1e0626d5bdb6d9120cdceed393c87b61d31ba972efb3e57ff1a9e856b1ed4e7d72b8b46082cdcde14bc4d6eb99fac74dfb304d2cd@34.82.221.18",
		"enode://584b5500f405eae3204714becab268305beee5418011cbfdea484ab158721ca8acba135120e573e00e26a1962d9ce2b1e6236763c82528c6639e381b2cfcb4bd@35.184.84.143",
		"enode://e22017240b760e95065d011f4e0017bb5978bd25b039846faa52bf18e6a2015d03fdd3f62ced12222af8bddc9324a5eda3ca31beaf1d1b839cdd63f887cb907f@35.196.20.106",
		"enode://f59e7b9f4fceeca859346be234297bf91c0d8412f6dcc0a137614ab9c13feaf0bb7e281772c7cc2c2f426edeab7b83d35c16984caf005d5abafb1d8ee086ecd3@35.202.3.128",
		"enode://efa8fda719349fed7e3a253e6d38c48ae60fda054f8f75acabdc619856c86a64112d63fc563f9e8aaa11ea47ed777d071d7060439d1ffe9082bfe541e45e81a8@35.224.101.46",
		"enode://ca5b1882b2231adc861b973066a8117904e9302fb9a93a2acf07fcb13029e5555e1b26844734e43cf5030c0ff7277f65294f96e2e294f25ff2882a25d9da1d47@35.231.200.19",
	},
	"cyprus1": {
		"enode://f0167ae476d5abc03f763f227ac4867c82ef0f52ebeb23f2af91bc6f60caed5d4ab6c5e04d543b2b5fb717d5cb169d6999a6e9b4b86834ec10e296d57027cc73@34.139.206.86",
		"enode://72681d518f4acb29e2635a1129eb4276796ff84700d4bc5b20ab5d7b294c22f9791c9529f09e8eb061e6a29244bf7e905e205c717481078ec6ca1e30bb2facdc@34.168.1.153",
		"enode://ab4b68edc8c743ad11c3c999b281e18d3ad7f2ce7a8575af88b8dc4bd428ae374bf7bdb7fbeafb235001eff3b887c619a1db1fcc5ad0c2396ce5cebf09efe4fb@34.168.35.191",
		"enode://7105a4bb6a79af980e7caa15d447b40967d47e19bc5236659230790a115c36e4d2311bfedd97eba059cd4e30e7e61d0b66a42fecb1d164d844a1e19985e4a7be@34.82.221.18",
		"enode://f6ae236c7d7b741df63d3f74164e98cdf5592d1075cb446534422a0d27bcba37faab34e107b768224f1d3d60a9bed5e60933fe3ee5431b3a2a44aba212217335@35.184.84.143",
		"enode://a08d1c79296eb512e3193df47bb357ff95cdaa8fe588471a9c3f05ba17845fbeb0b27e08eed1e1938e8f60f2c38c1cc7e93aecbac4de0471cc185280627d81e8@35.196.20.106",
		"enode://c70481de8f2aca8b20292263764b80e2f83e82e6a280843c40497c56dd9f76036efe4f1cf0f1813660f11a3937da32f4b42f887aef5a398b8012680a8e01f63f@35.202.3.128",
		"enode://ab5e84b6b03c72bc7fe74b891946e36b1a3b0dd0732e0c56ed13050778d1f5fbc691a504d81e35fe5f80c0c63260719e5ebe8cf3acff2f3bfe354ec95c3edf5a@35.224.101.46",
		"enode://f2765c39ad77b4e544f63ed22d190c0e67142d0ff610c5e0e647e812adc2ffdfc2001f70881c5759155b2ecb254c43bb045b8b1c9f70e1fa8b5568b8ce906f30@35.231.200.19",
	},
	"cyprus2": {
		"enode://ddd8170b45829579694600543002d6cecd96538a9e45bcf97d20f8b18c2ee5b1b4f0e86eb00ca7d39f3c0b1f18b9b9b08b70f8d1e5bd8ee58595a15041c7f2e3@34.139.206.86",
		"enode://ddc70b9cdacb37d17e5d2efbb16a277c59a5a95dd0ac7a51133b2fd9c3a8669aa39fc7624bf88cd8531b88e42c876b3d125f71b1539a632f32a6c52a4849a41b@34.168.1.153",
		"enode://7df114e99278d1330e109d6ae08926e2b62135e0ce8605c9ad0cb2fda7a60ad15658dd7aaa1b51bcc0305546d5f94167d60068e95eb46bb639565fe3f5239314@34.168.35.191",
		"enode://0f4e2f0c2db5855da9191612ee624dd34d95fdb4c34bef3028721b7e60f84bf2b6cd7ff99658fa2a5e952e38b33236f8e44c9572c87a96972e56980bfc186a17@34.82.221.18",
		"enode://3f2488650b2bf6fa81d52361f7481f83e4c8af3c2a701f0b62043ed21ff9fb29019b05f9f0e825af18ffd4498b863aeac2bfdbc999ebbb10ae3e9804a8d4ec60@35.184.84.143",
		"enode://2d8bd796457fab15c1a8eec41167994e427b2adab19faca559ccb99f1db02cc6c8f3d79934a4c1222bc3dcaf5f79a5597ccb201f2776de76a3c69807f25f667c@35.196.20.106",
		"enode://019eb30aefd559428f14d0e52dd2c152b693af090550d99fa36d5e994faeeaa25589d5e7c47526f9c59b203a43afc9819bd77163fc70ee6062e1aa5b2b84a010@35.202.3.128",
		"enode://6669310a34abfee90286324cfc5f897666c3102521440cc65d3ec7cbda58429f1800b9815ffbc3f627f50095bad066eade772ec004a280c2971b9d0d2efffdf5@35.224.101.46",
		"enode://ae6b77c47515d4b115c093624af53814e29a8435403628c49b5b4709abf10f9a04473365ae555ff6979425620f526026d968b5ed9d228bcf1413d0dfa7a03f2e@35.231.200.19",
	},
	"cyprus3": {
		"enode://d7a04846b19ee4d23a3e71399147cb819b3798f275c95c3512aa6505555c5dd77d8c3dc99f725c893ede8ef1a42a2aa4d7d0bb30947b5ec1f3801d0f4fd1005d@34.139.206.86",
		"enode://30dd68756309ea4ed8a00e788c25a8cef85dab47d59fc52d7192fc33a69cef882fdfa26a60e376a0ede4b357dd063ea5e34d42b2d4265cc39ef0fb70a9abe818@34.168.1.153",
		"enode://ead01a774f4f786adfb9333faff1e3400909f17c5d222c4858f35f826cf9c73b6e7526c187d48631499034caadf5084e89f86a99f3d4f447f946c4fe1d855d91@34.168.35.191",
		"enode://b94e17336d7b781ded8c99f67c80363be4437cade8103fc8e0cb31075f8efbf04ac186175f3ee38ec889b3b90d178735eefbca909a39cd22eeb0f674ea922fba@34.82.221.18",
		"enode://d23321231b7d74f4e9ffac71e78167cbe6cf0ffc088893146ea3803437562e0b24cd7f9ea869da4da10297649a11bf7256a8a90fea7ac343c4d35147aad7c0dc@35.184.84.143",
		"enode://52eb6b9f9dace4203ebca65cbe32dd4003ac0a8cb35cdb99edc6a7ab1365f68b697824ab39a9b1686267aba1b469761d4eedb78fcd395520a3b78cc760761f2b@35.196.20.106",
		"enode://49063d167e256a67d2320ca8266ec388db5448bda85572ef4596a95f76e9cfa4040ad4bda9ecdb5bfe88f4a1d1d1c718be63e67592c69d14e722faa4b5d78759@35.202.3.128",
		"enode://c7b5f3577b0b7382bd680f9abb394cc786a3abca1a922a5598f1370b4e23d2ad0197c26e63a0a51402827ab239852807889eee02d718f41d78c8da478d97a02f@35.224.101.46",
		"enode://b21187460ac9afe2c99076acce6c2c7614ba8d5498cdb88afb7aafbdbf62e75ea9783025f3d2da2149c9c0292e6caaa3c2fe43a7d159aff40a5d7f4d95d80944@35.231.200.19",
	},
	"paxos1": {
		"enode://98c9db3d64fad4b8a10a0e2f07e248a4ecf1c5317753f761c2e1677b318ef4c029471ced18cbc8cc74808fbb22b9197ab9ae2f3ea65416b606769373a83c9b59@34.139.206.86",
		"enode://25b8ae067bd0089010b982da3296bac795f21eec56ed71f2c357459aee070e0b2d83c1823d63c1c381ab877b2000aab00fb0e276039cbe5adbd999b99bb55414@34.168.1.153",
		"enode://d8e6a4615c5cec9f0b74b586fd2f822a2b0f343e4614f7e05126f052b75e2e56abe866a81d8848a9ff7c7ab8d8cf08f14ca7b3a5384111001d23a7dc95c119b2@34.168.35.191",
		"enode://56a35e2fa91d0a1759e4582fd1dd514a82329dad959c81938177b93a12f844e62a39fe64c63f5fb3f69986a51d7c5d8916fd013a8a9e80c1a69ba834b69e4ee2@34.82.221.18",
		"enode://17a83dbaf39e2c311a41fb4dd69ff001bac68cd794a24d9108e5de77519d9aad41e5c6196c9a0af789980897d73f2674d555c34b12ad0846d2c145a3d64c2566@35.184.84.143",
		"enode://2fcdd3539d2ae834c8e4d128bf3500b0ac08fac20e322475986efa4213c470c96c29f508425d9d79830b536807801b8e61d8e36772a34e3b0da62d8c2da7b12b@35.196.20.106",
		"enode://6762182a7db2c352778f1f07e5bbc4f06a1b1f201beee93c6fd84b3063d5eae34850355fd61f72727571d466ae404a8ac805c157e26dd1791c4b3df01db7df1a@35.202.3.128",
		"enode://d2da1d562522383ed0bbe731218ed9319492f8d6abedf3ebe6b702da0b8ad679e77a0869851c4b6d12b954553ab33adf8440c86606a9e8bfa4b7aced37afe5aa@35.224.101.46",
		"enode://415954e25c89c0b2f90286dd0007cccf919091bd0a758cd8d3673aad3cf5cb2379ac4896f92bcb5c23b241ec7ac646f2e7d5266f56cfb5d05d3bebdd83570d6a@35.231.200.19",
	},
	"paxos2": {
		"enode://489fdf851aebcb4608f3b14633919a4b060cbdbad770566a13355c50a3ee9a23d7736c4eb8baf57a4ced53857bfe080ea19c8e9657717f82fff8ab9d41918b9b@34.139.206.86",
		"enode://9f811c7f3a96fe7eb956036d34fabace4e61f779db509b1a18af087e0f9324dc73672de69bec58170474d0842b5919f1cda0ce5195a5f9b977a2e948dd5724cb@34.168.1.153",
		"enode://d9c5ea4cc3a2435e989ebd37a7673f54cbc8b4f8b08c36705493d3fe6d473c38ce7d074646c7f5fe3be68381fb2950f94ebbe81278c8ef7e5d3a661a8cc70043@34.168.35.191",
		"enode://315015cda55d895c57ecb6cceddd455429f18917a52030b9a3b8048c1f9f65dbaaebfee84e8194c7b613a924e7433e2f1ae186011708b69ac8eeb7f39d8b0d04@34.82.221.18",
		"enode://95f08d5c79879176955251b510ba47f476fdb8d1260d34298d9976ed615c23074f97213f92fca0025181a4c050cefd06bbed61d777a82bd4b23235093f0c0818@35.184.84.143",
		"enode://e8f182f36f691715a384ae5cee7e0450b82321a8caaf56d2ac41eb19f4a818397d75aa658deb73eb426bc850a316bdff1c9ae2cfef521e6bc643d80882ad760d@35.196.20.106",
		"enode://b44da03d1a3c33ebd4309b8426f9de0637b24f097630ca542327d7f757428d954102d88c9850af07b2188d5d3ea138834facc11617588807c0ad8a734e35159a@35.202.3.128",
		"enode://e222f6f453ed0270cf6c481bbfa73751b7b8cc79e5fe46d9289955eab34319d4cbb29d0050fd84038d62518da9b289c81dafec9d668c3ae4700128c345867caf@35.224.101.46",
		"enode://374e25059c548c0280f667dc9b47330c567d1fc52e09c2ec7f4d8f236eedfee7207964f3824f1864ad4599744e8a9e1faff32a3090b35f276d7c30200d319818@35.231.200.19",
	},
	"paxos3": {
		"enode://fae7e4fe231032eb4ac4fd343e4015ec739d04a4dd1616a3b2a0ce40aefdc832c0f5803071124f6b484ce00cc68fb5752daa9535e429bf16aa2410f05642a31f@34.139.206.86",
		"enode://1f66bc6eb44da24338c4064a4f0911c9583034be54206b284bb248dcc8c2034e3f2ee3e114489953a89cc714bdd6fb25020fa719056f7cdfb4e0f034f3f1dc05@34.168.1.153",
		"enode://66ed87e05c82cebf5c7c3164f8ed5304fab5670848a3da9519bf1b64052138b68e974e8d1f2460b8e116882d38b03ac7277d74f36a63e661ed27a96393e235c7@34.168.35.191",
		"enode://9c241135c2badb75024b2f863c65afaa7539c5989cd362eb25a772aeba64649871f62a12b7c5adfc3208f4fc5c08c8398ea74af478b8ff2592bbb4861c99b671@34.82.221.18",
		"enode://dd09bebd672a6bd4e00af62805ca497cc4393dd082013655c61e79f0eba0979f76aa05920da9b5d7166c302db94819b7ca7d6c37b3de74519ee4457fd31f464f@35.184.84.143",
		"enode://6de68477e6b5f9b7e2bb8b525885823508dd8d70af99a27d7c3ff90f80d334935bdbfd05893131651726e4ed10b33b10ef4214f9f7a9b571037c2dd6cbe21348@35.196.20.106",
		"enode://6b1d51c6da03e7ad96dbbcc53fca9e70d4e6e270f423da4fe7c391c93ef5afd37f37c6853d39edb3878619683b8b732978ea947b682602547303c1f82670d6e9@35.202.3.128",
		"enode://a5f5d3b3e103042bcab82307d87bc6934c364e9f109303f719bf9ca199a56e0793cf1a4e81d820fee697f3a49ab9aba718bc5b0577ee710d12d82a63f68750e9@35.224.101.46",
		"enode://74676b92aef21c4acb9702f500bc018852281b461c02859995cb15af1704cee56012fb1dfca20ff8bbcb620666581f5882646d1bdec35bf72613c6dea4ab16eb@35.231.200.19",
	},
	"hydra1": {
		"enode://b817ff48c36a64bd007cefff0960c1f2b1bb0aed12b782d32746268af460f4ca6b0c67cc8f50669737d3be8a43db5040c39d6b3d47abbcc39204a184f2d4197a@34.139.206.86",
		"enode://eaf7485ab5a142425e35c98b375fa6778c54bdaabd611e2b55b75a25de4321bc8e5147cdca8b5c17e6761a7ac3be0628796f7559030552c11ef1f0be6a01ea27@34.168.1.153",
		"enode://18cd16ab91c763777c0b64d288831ea3d4e253c1af37fc88b8c5883e61f30a0cc28f3751ffa5d8cd8196608f60fa9200175faff5a9123200c18ee6e41b7a08e8@34.168.35.191",
		"enode://ef34655ff1179a58c281b8307e30f0e21d2dacfe6c7494fc6a6239f23d30980c4f0361e438db74e894228ec57af6efd392e20e6ab9842faa1842dcecaea56171@34.82.221.18",
		"enode://740f12afb9bd0b9977a7115b4eee4f4b60c836f49f09db870e725ba69225c88ffabb6fdbbb12420f3be999efc322d0fd482b4fc9da1c29b335c99993b9a0b817@35.184.84.143",
		"enode://adebacc6a682347b6b336aebebcb441b4ae9dd0487d6771f19f42efa89842abfa3cc2becebb5ea71f3f7f653fa2b806a40a9baac24ea2da5030617215e03e26d@35.196.20.106",
		"enode://a4ec6d75718ceeb7988bcbe31cd27f9ae60f0dba3be4e8bf810fee35d857eaaa17c9c65c2e4105b2828b9ab0155027e3084ddc819bcc6f3a71c78965df204844@35.202.3.128",
		"enode://a24f74921e581ff97101b69d2217326cd017320abd4f677fcba36590a76e8bc9dd6d7fb3b48c34091997c7d1e5279abb34cffcfba4b760050c3664d8e0207635@35.224.101.46",
		"enode://bbc7d7cb9c820cc50e131516b18fda1c8185bb63740ffd85e25936bd46838f1f68bb1c630ca3f86f0f6b474d6f53e128285227d0afd64dbffee1e9a0158d0f0d@35.231.200.19",
	},
	"hydra2": {
		"enode://70957213c38156cd2c97334e2f3cf22db3facc00b619771f0d4895070bf3d5f3662d642b5280c36a4920dc4236f79404e738e0cb0933a7b9693b171d394052d9@34.139.206.86",
		"enode://10d2ff6b2bba7284c9797b5676aa30c869b6e8490ed16ca0bb9d60bc149fca8cadb5227e30a9e9c6bd806cec9fd8c35491fa826b967a1d5eef6cb961bddc3847@34.168.1.153",
		"enode://8f8b9cd6b686069556e82c262a55d957edaae3f2b039ad60a533b8e2d1e514ad038020f946de9c425963c6a091f87d7afdd75fa743e079ac18e16273943526ee@34.168.35.191",
		"enode://e86bc13ee420e4896aa79ebd5244dc0d0c8cf9da00aa2fa1301ae80a131e3aa230b87f9111c1e6df181e885d2b712423b522cbfa58fca0812f90dd8613fdf7a9@34.82.221.18",
		"enode://4e1be25606cec5be210bab0c8e1ffb4b3ae6efb2f8e83f8158a4b60c7f4988ece1e784fea0ca12926c629594789e8cb84837dfac81b7daa390dfda73a7a754ea@35.184.84.143",
		"enode://e327b4aeb7ae97450b92a5a801feaaa0c7bdc0f19f4b373e69ebe1701b91c28c92ea5c012c9262038bda85e2064b0395dc28d73a187a73d3f197ffca82505d96@35.196.20.106",
		"enode://7df9332a19a277a957ff96117366113b7e66bbed1411b9174c1dd23860ca89f3de56b05bda5620129da21a9c8ce0857235b48dc5fe2d03e5e89af5c064095d49@35.202.3.128",
		"enode://774fc0c52ef53d44b2407c6edef1ce1291d72e33aeed1f1a82f6a6e289c1c1a936f9a100aba68ba28a63911dd682a731d67b6f539fbeacdbb2efda1305a1b838@35.224.101.46",
		"enode://2a7eea6c5a8a89aa8480f42a9c96b2e63fbbdc3acd8b94d991f75bea8b6b446ee11028903cb7fe8ca7d66cefc965338ec12f81b450655ac90e5ec8b3cb6123d3@35.231.200.19",
	},
	"hydra3": {
		"enode://f5c1ed2e7c88a53537cdd5fbaacb3d4cfa1d7c82a10575806fc60df54d58e2f7dd37dce001e595284193b115e486dd23ea66c1a6ac6c463dae905e8a97b4e156@34.139.206.86",
		"enode://3efc3b47aa3751ef238697ae6f6d764a6502d234640b7e93000ae8ab3755d1b72e5baed3849d2526db86abf94bf32c89bd2f9ca352156e6565fed5870dce5e35@34.168.1.153",
		"enode://b9a760f8d94f7e7c3380e1aab946213adc17970f5c79dbb035fd955c58ebd5e5050fe11aecbac2f5497e427de684b5494d51afd397c9a606f4b914fb075a4263@34.168.35.191",
		"enode://c0e69729c62f23a2bed0a57b15ae1c9db4e6e6639b53a5ff5310f3fa7134b1f284a50277f86c9f198253ddfceeaf12c857f619b057d1c49c8caaf9df9051801d@34.82.221.18",
		"enode://21f41b6cb952c6c7ef6e21b8a589abfa417252e30341be39002bebf6cd2d7da04f461bed117d73289955b999ece28c936a3e965d6d35957d97484049266f8708@35.184.84.143",
		"enode://75b9745000c1b30f82406cacaf449a3bb93679ef247b41c446c6bae8f00a64f5099c607a565543c432d2e1abc0d8b1649b666382cfd8891850c3844cf505f2b8@35.196.20.106",
		"enode://367f025a7df0ba1d08fa3b244dd2feba2718cd6e8305e36cdff993f3bea6bfa0f44d4c1382b9432a032b9d82ef51522b649a3340f1eff0d816748b06c313dd5f@35.202.3.128",
		"enode://3112f0b00fd3dbc7cc61a343559e42b68b0ac6222ae852ac945bb0d0054e00a7644a899700df0d37d8b8b179939ac409486b0a5532c3630fd0c59ef833d19c3f@35.224.101.46",
		"enode://a319cce2f164e90183757e3b0f1b7338f5d8646befb277fcf1a9afa2b98717118cea0993b056dd7c70f527f629cddf408ceda21f1d9f5d5f84e245f5c4905ab0@35.231.200.19",
	},
}

// GardenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Garden test network.
var GardenBootnodes = map[string][]string{
	"prime": {
        "enode://0ff1977bf3c64584d90ca5c29860542ce4c8a676c4566174f0a14136ec6fdbccd29984894d7d4c517fb10ae051c0a7785548d23a2d5271efc8d48cc56e6f852f@34.133.84.195:30303",
        "enode://56b5847709cbf199e4fa5f477eb77633936a5b13a9abe2ac5a2759fb7a99aeb4d5cdad9b4d2fdd91674e4c8c52f45580d0748ba07049c83a38b0ed8f561b72e1@35.238.43.34:30303",
        "enode://df53f9da03fde95b4d14988e362bc8921f718f94287e6fe2515bed63b258c02f1a4e22e5e5852faded48e7993fefcc4104419ea76acac00e2f09b8373b388e66@34.136.202.179:30303",
	},
	"cyprus": {
        "enode://12e47289e898504e257dd1b63c6913ccf54fe9440573b8e3804233e41b392596bc2c4973f5c44492e08a7695c6bea71d12275a95e3909e6409df295b4e2ea23a@34.136.202.179:30304",
        "enode://1758ade43c9e327c714a8987dd4e4ab413c2aa78d96043c07fb7dd699eff699171cca65da0b16d2b87d8dfc001c2fb0deed220971c6b984bc4fc542980314519@35.238.43.34:30304",
        "enode://c3f47efb75099e40602010e713de48a744540f9ef133774f6cec82b1fec02a5bb3b6f59ba639af2648986618e969499e4a98312fe087026e61d1b6f8cd8874c8@34.133.84.195:30304",
	},
	"paxos": {
        "enode://8b86826dfd26dd9c17321dfcf11bb202104258387b7cdb63d8363ee9a6101ae719a7c545270877c540bb08ffdf106b68ea8a7ec46177ed729ba2ef337348d25b@35.238.43.34:30305",
        "enode://b94ad87e43876ec455ade07a3065864d89308539caa21b77c6d935bc44fc3f9509b976f621af1bf0dd607feace711efd235303159ada39d5cdfe912aea7d7d24@34.133.84.195:30305",
        "enode://ca3638c5477ed4a521fa5f9f63934feb978a981c6306fbf89540c7e3f68587d085bcb0edf24cb649acfba220be56b3b441c1db5c1520e24d621395a77c36da87@34.136.202.179:30305",
	},
	"hydra": {
        "enode://5c0f9fcd636dd15f164df43e96268a48290761e9ff74675457442f1b25600d97c75ab471338fa6e6b28849beb87465f375a17b5dd3ff586b76a10f02a65a7858@35.238.43.34:30306",
        "enode://ce34485efac2d5b12e5b91d09ca37d27f3e1256753f46a25083ee8ba9a6169fa39821cd8830cfd6e5512c48f69507b072790739a75b6c3e1393bed6480e4f801@34.133.84.195:30306",
        "enode://ec73a93abd1dd8cd95ecb27466350fea8a570fd8637ddc7979227df69ea8198a2abc6b0499108c78b1026e6e216319be3391f8e2afcb7f7edfadf0bd318582dc@34.136.202.179:30306",
	},
	"cyprus1": {
        "enode://319d6cedcef315a1a30023f37273e050571523e1487873989f14ff6b57dd43eb11f467da4afeeadb81019ab51deda1a06f627e4f7fca5425f01835d59beea38d@35.238.43.34:30307",
        "enode://6d76f9abee47c622410fabd0a9be946299623dfd9313a1d0094b4cbd84814db8b3d339a83f9be9a33a2248479f8fda501a518ef09a7e2d4200908a1e6a8886cb@34.136.202.179:30307",
        "enode://8f979ae101fb6928d68f8ef071d6847b7bd733ea804752595137ff30c0b28e5832165d3fbf780f1b70b0fb3955fa387fac0f6a04dd5b75f4e8c4c8b0c628449e@34.133.84.195:30307",
	},
	"cyprus2": {
        "enode://364ae711b9b13b344e9f1b31ef569bc421df3e476216077d27d5b781104990f4d75f51b204dee418e20dc5047b3276554ce27d5c70d8277aa8446065f6a1dd96@35.238.43.34:30308",
        "enode://39882acf0d3bc460f0dd046cf84109192e23bbaac951c3e128e7bf73a3218eeb825903eed6b57491892f8376f741a9fc394fa83f482f0ea89df78c40f333ad49@34.136.202.179:30308",
        "enode://ddba4c858452ff777985bae6518d33ae1c2cf1fb72a6a3b43dea194b7d0ed859ed3e0c4d6dded642df58bec9a2ae45817ef70e23c6936da6a2121df1f18dc183@34.133.84.195:30308",
	},
	"cyprus3": {
        "enode://1d1115bbfb9852eac8312528387233fe59012c9f1dd92e8a5ef51cf5fd6c3eadbfcc027ddf757eb82219b2bf00aff340a731e7b7025ba6238e371a0fcc01366d@34.136.202.179:30309",
        "enode://3922525d55bceb76750efcc44120a595e2267a22d48499747545e023ecb17188b5c95149a999cc572fea0e09097e0803d7c74d465025791a8d7a6d8fff408521@35.238.43.34:30309",
        "enode://4e405e3befa076275fc8b7ba038eecd0aef2a0faf6b422e28adfa53367d306d1f14359c54bf58bc0812a43015c98c3948f93bc8e9ab77e74ae3977615b01b670@34.133.84.195:30309",
	},
	"paxos1": {
        "enode://6504d5b8d58b98c90ad9b5908f74f6c4779f6afd1102102e73623863b4a4a9fa2db8e8e24b36090aa90eb5aaf166486293749d20175bd5fdf417bcf10a1b015f@35.238.43.34:30310",
        "enode://743df6de87138b9dae975b0920351fe2a3bc75b3c7a545e60372ae697fa8675f8f598b3f54b4595b4702dc16497023130883d4e98200c7b7e795ca2d8f8b1435@34.133.84.195:30310",
        "enode://db24e8572789344e62fd836302340d911d4e3c532804bf31e8a2a5a66f891ae220fc425aba602a2e2bd9fca7a966793012854abad15367a88a2b2227ca7c2277@34.136.202.179:30310",
	},
	"paxos2": {
        "enode://724d5422bd2fe1bc397aef9fc20b3ece83a90f66b59725e43ae4c5e50830f85d364b500fa49d9116fa4fcdb72826507457f5fca4b241d358d551ff9ee8c19de7@35.238.43.34:30311",
        "enode://95f90aa3a65a7eee159f09495a4f2ba75ff55ad7e305ce591ff69ad4406ae8c3b704ae9b9828666954810e203ed6ee6e90893c9941f5db4734b5787891635adc@34.136.202.179:30311",
        "enode://97a86eb8e2dce83d8956762bbd916a382361b1e83ae449a44050db0f8bda641ab8bd4932bff2e2d0ea9be8b30220269c93ef0436c26d0f1662420b34f640422d@34.133.84.195:30311",
	},
	"paxos3": {
        "enode://5975da15e0e3a0931dfaf72870187eb66af867f21d4fcbc98c84338f7fe10bc40e9fae90d1b0fe1166567288669e0f64b3da13e22c494b072d151ce10bfff518@34.133.84.195:30312",
        "enode://81e5c8a67f326c34c49e97a48986bbacb15eb7ebce20b8d7228ed93e3ed82f9d2b99dd03cea203bdc3f84e17b82399eb04e7ebd47eba02f2c82a044afce75c8d@34.136.202.179:30312",
        "enode://9bf904b9b11dc76a5b1201d54d844657e7e2709754741f5232ef4749d7782ba699266f9eec661e554fb5744f73ee4bc210360f1b64aa05f52c6a093f79625618@35.238.43.34:30312",
	},
	"hydra1": {
        "enode://707aae3ca6fbdd0b1b0d81b4ef4a74c534f23d54fdb4f5c57e826b12fbc247b787ac28e28b958d41b556e93773f4d94db7b624cb7043a0a8e9f969c0fd07794d@35.238.43.34:30313",
        "enode://a5a2a352db4d15a3d0e824072b8fdd0b792a9eaeafd5cb35c43b934a23f6dfbf3d6dfd960a63ec886201cdfff1059105d2b302c216e3def6209de9290925028d@34.133.84.195:30313",
        "enode://e1d2240dfe144ecb521aa924e744fd3c622f92e7edabbbb7bb37f1f6e7902a30df9a6f27c2b0b95cc6cce08a79f629c83123ccf2225a6cdc52ada74e12f201cb@34.136.202.179:30313",
	},
	"hydra2": {
        "enode://396d0de08044ab1e6a546b613a3f43a376435a1cfbb44f3d1dab1a91dff576abc7b0c18ceed74afadaf1980547d1c67ad9574321fa45fdc6245325310c35b111@34.133.84.195:30314",
        "enode://9691f6644f5d67f74463ce3f954048a6bf5092c76c06a3cf8175ef9d0c9987141f5f33be857f7b82f594115e3871114ee86578fd729eb6ee329a2b699bc87602@34.136.202.179:30314",
        "enode://b9254fc3358f4b6113ed3422e5a17d91febbe1e25714fa28f68d619a42976af68b4f0c94d87e3365a5b0178c564f5492846201cb4f350065fdb8eb2dda3ea600@35.238.43.34:30314",
	},
	"hydra3": {
        "enode://2803d0c9d7dfad8a7a609aeb24b15e8b5245aecb625f37032413a0ebc314172d88921efd055a5ce838085aae17337d2db78dfc3b463ffdac8fbe49b8088b441c@34.133.84.195:30315",
        "enode://35743013ef3f61b7e2cdf0385447a87723dbeb850317de399c91a532d07e29e2a56ce5d08b1c7bfe876a649d8742b1a4a2470bcb1b6b7f7492da40ecac68188e@34.136.202.179:30315",
        "enode://53d2cd44c5fd7ff400cdd03ab7614ec07b96e9c1851cfa09606df42d61d0a8f6a520f7480750fad6873431098922c31ad03e85a4b6a6883ed70bc95354a91619@35.238.43.34:30315",
	},
}


// OrchardBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Orchard test network
var OrchardBootnodes = []string{
	"enode://f99403abcfbee37e3232e6bb4d7fda4d70496585a53857ccb4aff5ec32d0f27186b5097430d5806f20f2003f35cfec5d778598a3945d530f212b7072caab9b8a@35.188.17.207",  // us-central1-b
	"enode://142e48e3c36e5fe21aebf2941f2e63eb4720febe67de17dd84baf010e33c76275567ede53674007ab2848eec53022cd0cb94bcbea10960ae93edb5497c8caa2a@104.198.69.162", // us-central1-a
	"enode://d6d27b273682f8abc7ffff04dc9006bd40f0a079a8ba24da351b714506bb82c1f106ff073fa04983345aef15c876c602209b48b37d8ee10dad581fd1d9db9263@34.23.150.43",   // us-east1-c
}

// LighthouseBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Lighthouse test network
var LighthouseBootnodes = []string{
	"enode://ee89c22bff79d040fcf3dbaea3bcbe429e68b0ca9e32671027554e96aea3f132f6abed8cf5be514c50b76cf2cab96c7d9064a93f0bbd0903f26df4be01ce0e6a@35.196.124.28", // europe-southwest1-a
	"enode://b39cf3080c8c9165bf0b50a7f6c8ff5a3568649b0c57ae786f630a054722fccfec7e3232594eb37a62a04e7c310a4d4e899ea42c0bd5a5043a248510715e2af9@35.187.55.110", // southamerica-east1-b
	"enode://ce3daf05c462b36bc1a6261b8edb0cd72bf041c1f8fb59100d6a09dc3415d1f136cc8768b12129db73c48085c8e01d7a606be447d3f681f405ab427641599235@34.92.50.205",  // asia-northeast3-a
}

var V5Bootnodes = []string{}

const dnsPrefix = "enrtree://ALE24Z2TEZV2XK46RXVB6IIN5HB5WTI4F4SMAVLYCAQIUPU53RSUU@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case ProgpowColosseumGenesisHash:
		net = "colosseum"
	case ProgpowGardenGenesisHash:
		net = "garden"
	case ProgpowOrchardGenesisHash:
		net = "orchard"
	case ProgpowLighthouseGenesisHash:
		net = "lighthouse"
	case Blake3PowColosseumGenesisHash:
		net = "colosseum"
	case Blake3PowGardenGenesisHash:
		net = "garden"
	case Blake3PowOrchardGenesisHash:
		net = "orchard"
	case Blake3PowLighthouseGenesisHash:
		net = "lighthouse"
	default:
		return ""
	}
	return dnsPrefix + common.NodeLocation.Name() + "." + net + ".quainodes.io"
}
