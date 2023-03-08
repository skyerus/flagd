# Changelog

## [0.2.7](https://github.com/skyerus/flagd/compare/v0.4.1...v0.2.7) (2023-03-08)


### ⚠ BREAKING CHANGES

* Use OTel to export metrics (metric name changes) ([#419](https://github.com/skyerus/flagd/issues/419))
* consolidated configuration change events into one event ([#241](https://github.com/skyerus/flagd/issues/241))
* start command flag refactor ([#222](https://github.com/skyerus/flagd/issues/222))
* Updated service to use connect ([#163](https://github.com/skyerus/flagd/issues/163))

### ✨ New Features

* add debug logging for merge behaviour  ([#456](https://github.com/skyerus/flagd/issues/456)) ([dc71e84](https://github.com/skyerus/flagd/commit/dc71e84f0704690b528e7f1c2b56cb4898374fbf))
* add Health and Readiness probes ([#418](https://github.com/skyerus/flagd/issues/418)) ([7f2358c](https://github.com/skyerus/flagd/commit/7f2358ce207527c890f4a2f46ce4b9e8bf697095))
* Add version to startup message ([#430](https://github.com/skyerus/flagd/issues/430)) ([8daf613](https://github.com/skyerus/flagd/commit/8daf613e7e4f4492df0c06e2ef464f4337cadaca))
* consolidated configuration change events into one event ([#241](https://github.com/skyerus/flagd/issues/241)) ([f9684b8](https://github.com/skyerus/flagd/commit/f9684b858dfef40576e0031654b421a37e8bb1d6))
* enable request logging via the --debug flag ([#226](https://github.com/skyerus/flagd/issues/226)) ([11954b5](https://github.com/skyerus/flagd/commit/11954b521cc6197d0dc04b163e66e38d4c288047))
* Eventing ([#187](https://github.com/skyerus/flagd/issues/187)) ([3f7fcd2](https://github.com/skyerus/flagd/commit/3f7fcd2f57318fad4e0bf501f202af990d3c5a79))
* file extension detection ([#257](https://github.com/skyerus/flagd/issues/257)) ([ca22541](https://github.com/skyerus/flagd/commit/ca2254117adc163b94d662b3d1fbfd868f788fcb))
* fixing informer issues ([#191](https://github.com/skyerus/flagd/issues/191)) ([837b0c6](https://github.com/skyerus/flagd/commit/837b0c673e7e7d4799f100291ca520d22944f22a))
* flagd image signing ([#338](https://github.com/skyerus/flagd/issues/338)) ([eca6a60](https://github.com/skyerus/flagd/commit/eca6a60967999a303ceef5465f1acc35c83afd6d))
* introduce flag merge behaviour ([#414](https://github.com/skyerus/flagd/issues/414)) ([524f65e](https://github.com/skyerus/flagd/commit/524f65ea7215466bb4ac24a8d0d5953dd1cfe9a0))
* introduce grpc sync for flagd ([#297](https://github.com/skyerus/flagd/issues/297)) ([33413f2](https://github.com/skyerus/flagd/commit/33413f25882a3f1cf4953da0f18e746bfb69faf4))
* only fire modify event when FeatureFlagConfiguration Generation field has changed ([#167](https://github.com/skyerus/flagd/issues/167)) ([e2fc7ee](https://github.com/skyerus/flagd/commit/e2fc7ee2570a119923bf95b40b2046dfa4705f20))
* refactor and improve K8s sync provider ([#443](https://github.com/skyerus/flagd/issues/443)) ([4c03bfc](https://github.com/skyerus/flagd/commit/4c03bfc812e7ceabcac0979290bd74d9efc9da15))
* ResolveAll endpoint for bulk evaluation ([#239](https://github.com/skyerus/flagd/issues/239)) ([6437c43](https://github.com/skyerus/flagd/commit/6437c43022b5c94d2fb835a406d85a4e836f2fcf))
* Resurrected the STATIC flag reason. Documented the caching strategy. ([#224](https://github.com/skyerus/flagd/issues/224)) ([5830592](https://github.com/skyerus/flagd/commit/5830592053c55dc9e55c16854e40c3fc8345d6d1))
* snap ([#211](https://github.com/skyerus/flagd/issues/211)) ([c619844](https://github.com/skyerus/flagd/commit/c61984448d5cdadec62b5cf6f7e24fc5f75a3738))
* start command flag refactor ([#222](https://github.com/skyerus/flagd/issues/222)) ([14474cc](https://github.com/skyerus/flagd/commit/14474ccf65b9b92213e8c792e94c458022484df4))
* support yaml evaluator ([#206](https://github.com/skyerus/flagd/issues/206)) ([2dbace5](https://github.com/skyerus/flagd/commit/2dbace5b6bb8e187a7d44a3d3ec14190c63b3ae0))
* update in logging to console and Unify case usage, seperators and punctuation for logging ([#322](https://github.com/skyerus/flagd/issues/322)) ([0bdcfd2](https://github.com/skyerus/flagd/commit/0bdcfd2fecc03b15be9fc4b0489431b8fa86aed8))
* Updated service to use connect ([#163](https://github.com/skyerus/flagd/issues/163)) ([828d5c4](https://github.com/skyerus/flagd/commit/828d5c4c11157f5b7a77f5041806ba2523186764))
* Use OTel to export metrics (metric name changes) ([#419](https://github.com/skyerus/flagd/issues/419)) ([eb3982a](https://github.com/skyerus/flagd/commit/eb3982a1cb72d664022b5cb126b533cf61497001))


### 📚 Documentation

* add .net flagd provider ([73d7840](https://github.com/skyerus/flagd/commit/73d7840c9fdef9c62371c677e02c0d9773c85f95))
* configuration merge docs ([#455](https://github.com/skyerus/flagd/issues/455)) ([6cb66b1](https://github.com/skyerus/flagd/commit/6cb66b14d01b6ee1c270bbdd3e30d4016757eae5))
* documentation for creating a provider ([#413](https://github.com/skyerus/flagd/issues/413)) ([d0c099d](https://github.com/skyerus/flagd/commit/d0c099d9aba3ed4d760a1858381f5e29b6d49a9c))
* fixed typo ([#319](https://github.com/skyerus/flagd/issues/319)) ([1bfb321](https://github.com/skyerus/flagd/commit/1bfb321e3f9a8761ea06156e4b241a281cf3b533))
* Restructure flagd docs ([#231](https://github.com/skyerus/flagd/issues/231)) ([d614749](https://github.com/skyerus/flagd/commit/d6147490fd939f227025b2875b2a1888b4e4a3fe))
* Update docs ([#286](https://github.com/skyerus/flagd/issues/286)) ([702febc](https://github.com/skyerus/flagd/commit/702febc0dbc18b4934d8f08c1c2525c4d8a430ca))
* updated filepaths for schema store regex ([#344](https://github.com/skyerus/flagd/issues/344)) ([2d0e9d9](https://github.com/skyerus/flagd/commit/2d0e9d956fbc99f2775821cfecdceb2b016d2b78))


### 🔄 Refactoring

* implement delete event type in filepath sync, expand tests ([#369](https://github.com/skyerus/flagd/issues/369)) ([e906757](https://github.com/skyerus/flagd/commit/e906757a4bac6a556ed1e2a0cd7428d708d5caed))
* introduce add, update and delete operations to sync providers ([#307](https://github.com/skyerus/flagd/issues/307)) ([0c9a6c3](https://github.com/skyerus/flagd/commit/0c9a6c34e32bee8c7a8478d0844de7c9f6b4805f))
* isolate flag state management to store ([#383](https://github.com/skyerus/flagd/issues/383)) ([fef9b87](https://github.com/skyerus/flagd/commit/fef9b8756430f49df6622a768cc9d884c92ecd13))
* log clarification + cleanup ([#359](https://github.com/skyerus/flagd/issues/359)) ([541589b](https://github.com/skyerus/flagd/commit/541589b37ccdfba2e1257c6a29f16d2e2d05e7bf))
* logging standardisation ([#207](https://github.com/skyerus/flagd/issues/207)) ([418b9ef](https://github.com/skyerus/flagd/commit/418b9ef5210ba2f33800dabc9328d8eca2153927))
* refactor sync providers ([#291](https://github.com/skyerus/flagd/issues/291)) ([46c5bf8](https://github.com/skyerus/flagd/commit/46c5bf87267de8120a8de39c9e35f2ceb6f6034f))
* remove unused struct field ([#458](https://github.com/skyerus/flagd/issues/458)) ([a04c0b8](https://github.com/skyerus/flagd/commit/a04c0b837dbe9e28d1e01e43ea9e378a6c0f316a))
* simplify connect Resolve... funcs ([#391](https://github.com/skyerus/flagd/issues/391)) ([feebbe2](https://github.com/skyerus/flagd/commit/feebbe26e6b05e3b2f03c5e9a202fa899d814f68))
* Static reason type removal ([#165](https://github.com/skyerus/flagd/issues/165)) ([b655e88](https://github.com/skyerus/flagd/commit/b655e8848472679355b200f19f44b6dc55d0f0ed))


### 🐛 Bug Fixes

* ApiVersion check fix ([#193](https://github.com/skyerus/flagd/issues/193)) ([3a524d6](https://github.com/skyerus/flagd/commit/3a524d646187355bb224100f436c7b5f35abea3e))
* bubbles up unclean error exits ([#170](https://github.com/skyerus/flagd/issues/170)) ([9f7db02](https://github.com/skyerus/flagd/commit/9f7db0259d2d24cb880eeddaebd3b8f48758248a))
* changed eventing configuration mutex to rwmutex and added missing lock ([#220](https://github.com/skyerus/flagd/issues/220)) ([5bbef9e](https://github.com/skyerus/flagd/commit/5bbef9ea4b1960686e58298c2c2e192ca99f072f))
* checkout release tag before running container and binary releases ([#171](https://github.com/skyerus/flagd/issues/171)) ([50fe46f](https://github.com/skyerus/flagd/commit/50fe46fbbf120a0657c1df35b370cdc9051d0f31))
* CVE-2022-32149 ([#198](https://github.com/skyerus/flagd/issues/198)) ([11a7b34](https://github.com/skyerus/flagd/commit/11a7b3472ab2bc39bce7c40037e8f83736065163))
* deprecation warning fix ([#317](https://github.com/skyerus/flagd/issues/317)) ([a2630db](https://github.com/skyerus/flagd/commit/a2630dbba151f35cc985d38de9cf25bfee2b76c8))
* **deps:** update goreleaser/goreleaser-action action to v4 ([#340](https://github.com/skyerus/flagd/issues/340)) ([b9fcd5c](https://github.com/skyerus/flagd/commit/b9fcd5caa67a61b447c437b651471b4603b2b272))
* **deps:** update kubernetes packages to v0.26.1 ([#267](https://github.com/skyerus/flagd/issues/267)) ([26825f2](https://github.com/skyerus/flagd/commit/26825f288c56df638fd160caa93f926a8c136108))
* **deps:** update kubernetes packages to v0.26.2 ([#450](https://github.com/skyerus/flagd/issues/450)) ([2885227](https://github.com/skyerus/flagd/commit/28852270f34ff81c072337b29aa17f4b6634e9cc))
* **deps:** update module github.com/bufbuild/connect-go to v1.4.1 ([#268](https://github.com/skyerus/flagd/issues/268)) ([712d7dd](https://github.com/skyerus/flagd/commit/712d7dd4a34980bf9eddad99d926cbdd5d69d624))
* **deps:** update module github.com/bufbuild/connect-go to v1.5.0 ([#326](https://github.com/skyerus/flagd/issues/326)) ([7f332e5](https://github.com/skyerus/flagd/commit/7f332e50ecb1cea19108d1fa2fd79da3d5864bf9))
* **deps:** update module github.com/bufbuild/connect-go to v1.5.1 ([#365](https://github.com/skyerus/flagd/issues/365)) ([e25f452](https://github.com/skyerus/flagd/commit/e25f452906e034e339309270cc8db6dcd58e9973))
* **deps:** update module github.com/bufbuild/connect-go to v1.5.2 ([#416](https://github.com/skyerus/flagd/issues/416)) ([feb7f04](https://github.com/skyerus/flagd/commit/feb7f047365263758a63d8dffea936f621a4966d))
* **deps:** update module github.com/diegoholiveira/jsonlogic/v3 to v3.2.7 ([#283](https://github.com/skyerus/flagd/issues/283)) ([2ab5a00](https://github.com/skyerus/flagd/commit/2ab5a00fa6f19c7e0fe1a4e36649fae2633ac269))
* **deps:** update module github.com/mattn/go-colorable to v0.1.13 ([#260](https://github.com/skyerus/flagd/issues/260)) ([5b11504](https://github.com/skyerus/flagd/commit/5b11504cdce50c137540cc79d2db94e70a21338b))
* **deps:** update module github.com/open-feature/go-sdk-contrib/providers/flagd to v0.1.10 ([#459](https://github.com/skyerus/flagd/issues/459)) ([cbdf9b0](https://github.com/skyerus/flagd/commit/cbdf9b07c30239d7d04ef770cf4461fb33422fe9))
* **deps:** update module github.com/open-feature/go-sdk-contrib/providers/flagd to v0.1.9 ([#427](https://github.com/skyerus/flagd/issues/427)) ([42d2705](https://github.com/skyerus/flagd/commit/42d270558bf9badcff9a9b352fda35491c45aebe))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.23 ([#261](https://github.com/skyerus/flagd/issues/261)) ([a1dd3b9](https://github.com/skyerus/flagd/commit/a1dd3b9005374b5527f12b8e138250cacddc71af))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.24 ([#290](https://github.com/skyerus/flagd/issues/290)) ([38d3eba](https://github.com/skyerus/flagd/commit/38d3ebaffcb1f36a38003273c62c6317f0ee75a3))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.25 ([#324](https://github.com/skyerus/flagd/issues/324)) ([ed1d3aa](https://github.com/skyerus/flagd/commit/ed1d3aaba4ca179a89757a6b1c3f328826e787fc))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.26 ([#331](https://github.com/skyerus/flagd/issues/331)) ([be67e5f](https://github.com/skyerus/flagd/commit/be67e5f5bc1fb7351a04ffc4180447a27d57d32a))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.27 ([#335](https://github.com/skyerus/flagd/issues/335)) ([824cf1a](https://github.com/skyerus/flagd/commit/824cf1ab0f2e18826207af16d5328b817c755c8e))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.28 ([#342](https://github.com/skyerus/flagd/issues/342)) ([e6df80f](https://github.com/skyerus/flagd/commit/e6df80fd25d3da342e72d2ca0e923d9bf3d3f797))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.29 ([#429](https://github.com/skyerus/flagd/issues/429)) ([b7fae81](https://github.com/skyerus/flagd/commit/b7fae81b89b3a1a0793a688c32569c4284633c6a))
* **deps:** update module github.com/rs/cors to v1.8.3 ([#264](https://github.com/skyerus/flagd/issues/264)) ([0e6f2f3](https://github.com/skyerus/flagd/commit/0e6f2f3e5a77dae7d491eaf1094a65e692bebe5d))
* **deps:** update module github.com/spf13/viper to v1.15.0 ([#296](https://github.com/skyerus/flagd/issues/296)) ([d43220b](https://github.com/skyerus/flagd/commit/d43220b2be58e4bce05050c5d1b36788289ae7cc))
* **deps:** update module github.com/stretchr/testify to v1.8.1 ([#265](https://github.com/skyerus/flagd/issues/265)) ([2ec61c6](https://github.com/skyerus/flagd/commit/2ec61c6bc61c266451b496ff18c3dd9a74173233))
* **deps:** update module github.com/stretchr/testify to v1.8.2 ([#440](https://github.com/skyerus/flagd/issues/440)) ([ab3e674](https://github.com/skyerus/flagd/commit/ab3e6748abc7843c022afeaf7cb11193cdcf59c5))
* **deps:** update module golang.org/x/net to v0.6.0 ([#396](https://github.com/skyerus/flagd/issues/396)) ([beb7564](https://github.com/skyerus/flagd/commit/beb756470b1e1d5ef0670b8322b6ed9cb44efa24))
* **deps:** update module golang.org/x/net to v0.7.0 ([#410](https://github.com/skyerus/flagd/issues/410)) ([c6133b6](https://github.com/skyerus/flagd/commit/c6133b6af61f3d73ae73d318a1a9f44db2540540))
* **deps:** update module golang.org/x/net to v0.8.0 ([#468](https://github.com/skyerus/flagd/issues/468)) ([10d5f2c](https://github.com/skyerus/flagd/commit/10d5f2c55081a25daa1f0e0fa81f96f0fffbbc7b))
* **deps:** update module google.golang.org/grpc to v1.52.1 ([#314](https://github.com/skyerus/flagd/issues/314)) ([ad25388](https://github.com/skyerus/flagd/commit/ad25388461816100e19bda44a8e0077770ea0ee4))
* **deps:** update module google.golang.org/grpc to v1.52.3 ([#325](https://github.com/skyerus/flagd/issues/325)) ([8013ea5](https://github.com/skyerus/flagd/commit/8013ea5c6fa311b337c7ec1b1e8e756080808948))
* **deps:** update module google.golang.org/grpc to v1.53.0 ([#388](https://github.com/skyerus/flagd/issues/388)) ([174cd7c](https://github.com/skyerus/flagd/commit/174cd7c70fa5ae2573db2c5972b75786633e2f41))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.2 ([#336](https://github.com/skyerus/flagd/issues/336)) ([836d3cf](https://github.com/skyerus/flagd/commit/836d3cf3c06570d59929c3464e3c8e11c9b5a2fa))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.3 ([#372](https://github.com/skyerus/flagd/issues/372)) ([330ac91](https://github.com/skyerus/flagd/commit/330ac91e375124826b2a7a1a22d0daa18368ab99))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.4 ([#374](https://github.com/skyerus/flagd/issues/374)) ([d90e561](https://github.com/skyerus/flagd/commit/d90e561bfc5b798d13d4ba8f30f523b1053f3748))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.5 ([#454](https://github.com/skyerus/flagd/issues/454)) ([f907f11](https://github.com/skyerus/flagd/commit/f907f114f23fa2efa2637e254e829e4d53a90b51))
* error handling of Serve/ServeTLS funcs ([#397](https://github.com/skyerus/flagd/issues/397)) ([8923bf2](https://github.com/skyerus/flagd/commit/8923bf2d407e18b65c188aef9bf7370fc74c3be2))
* fix broken image signing ([#461](https://github.com/skyerus/flagd/issues/461)) ([05bb51c](https://github.com/skyerus/flagd/commit/05bb51c7ab30f6e976b87f54ca889e978f834211))
* fix race in http sync test ([#401](https://github.com/skyerus/flagd/issues/401)) ([1d0c8e1](https://github.com/skyerus/flagd/commit/1d0c8e168b73f7fbd4b27ece733041bbe08261c0))
* fix unbuffered channel blocking goroutine  ([#358](https://github.com/skyerus/flagd/issues/358)) ([4f1905a](https://github.com/skyerus/flagd/commit/4f1905a9ac6d62b5edb297fba904aac8680c89cf))
* fixing image delimeter  ([#463](https://github.com/skyerus/flagd/issues/463)) ([b4ee495](https://github.com/skyerus/flagd/commit/b4ee495dc8e00b032518ea42d272a36b3b662e95))
* improve invalid sync URI errror msg ([#252](https://github.com/skyerus/flagd/issues/252)) ([5939870](https://github.com/skyerus/flagd/commit/5939870b8994dbca585c53dd022485090aab2406))
* introduced RWMutex to flag state to prevent concurrent r/w of map ([#370](https://github.com/skyerus/flagd/issues/370)) ([93e356b](https://github.com/skyerus/flagd/commit/93e356b4ab0b65c71659bd52d73f618edffc96f5))
* omitempty targeting field in Flag structure ([#247](https://github.com/skyerus/flagd/issues/247)) ([3f406b5](https://github.com/skyerus/flagd/commit/3f406b53bda8b5beb8b0929da3802a0368c13151))
* remove non-error error log from parseFractionalEvaluationData ([#446](https://github.com/skyerus/flagd/issues/446)) ([34aca79](https://github.com/skyerus/flagd/commit/34aca79e6ec9876a6cced0fe49e1ceea34d83696))
* replace character slice with regex replace ([#250](https://github.com/skyerus/flagd/issues/250)) ([c92d101](https://github.com/skyerus/flagd/commit/c92d1012b0de6af694c3af2fede28053e2572b04))
* sbom artefact name ([#380](https://github.com/skyerus/flagd/issues/380)) ([3daef26](https://github.com/skyerus/flagd/commit/3daef263c43ed63776d604d27f7ae6b993fff143)), closes [#379](https://github.com/skyerus/flagd/issues/379)
* security issues ([#464](https://github.com/skyerus/flagd/issues/464)) ([7f1e759](https://github.com/skyerus/flagd/commit/7f1e759a87a9af63e9384005c959a3f500cc474c))
* send datasync on remove fs events ([#339](https://github.com/skyerus/flagd/issues/339)) ([4c9aaac](https://github.com/skyerus/flagd/commit/4c9aaaca77b1c8b16f59434aeb37407fced47ecf))
* set readiness once only ([#465](https://github.com/skyerus/flagd/issues/465)) ([41a888d](https://github.com/skyerus/flagd/commit/41a888d6b60c030b913280c2a1eeff8b25e8aada))
* set ResolveObject reason ([#375](https://github.com/skyerus/flagd/issues/375)) ([dcf199d](https://github.com/skyerus/flagd/commit/dcf199dab9d11b86454028869a54d77a474fc4a6))
* Update flagd systemd config to use URI ([#315](https://github.com/skyerus/flagd/issues/315)) ([93a04b4](https://github.com/skyerus/flagd/commit/93a04b46133e9220ec6f23d833c11f195e05c13e))
* update outdated doc link in deprecation warning ([#316](https://github.com/skyerus/flagd/issues/316)) ([19695d2](https://github.com/skyerus/flagd/commit/19695d2715129d6718ca0617b6ec6922ffb79c9b))
* updated merge functionality ([#182](https://github.com/skyerus/flagd/issues/182)) ([94d7697](https://github.com/skyerus/flagd/commit/94d7697d08a07cede4a548ef998792d00f8954a0))
* updated tcp listener ([#174](https://github.com/skyerus/flagd/issues/174)) ([b750ed1](https://github.com/skyerus/flagd/commit/b750ed1268b5e6efe779a34e764bad3e781f8e93))
* upgrade package containing vulnerability ([#162](https://github.com/skyerus/flagd/issues/162)) ([82278c7](https://github.com/skyerus/flagd/commit/82278c7cf08cc6b50f49ab500caf6f9003fc0823))
* use event.Has func for file change notification handling (increased stability across OS) ([#361](https://github.com/skyerus/flagd/issues/361)) ([09f74b9](https://github.com/skyerus/flagd/commit/09f74b9c5d15622c98da08558cbcd63fe9422754))


### 🧹 Chore

* add [@kavindu-dodan](https://github.com/kavindu-dodan) to CODEOWNERS ([#377](https://github.com/skyerus/flagd/issues/377)) ([e02c899](https://github.com/skyerus/flagd/commit/e02c8995d8329751b02d6c3016c33143e0b4ad4e))
* add additional sections to the release notes ([#449](https://github.com/skyerus/flagd/issues/449)) ([798f71a](https://github.com/skyerus/flagd/commit/798f71a92d2e2f450a53cda93b44217cbb2ad7fd))
* add architecture overview for recent refactoring ([#310](https://github.com/skyerus/flagd/issues/310)) ([a114070](https://github.com/skyerus/flagd/commit/a1140708503381ce3ddc99888f182a26d3842bfd))
* add sbom generation for releases ([#337](https://github.com/skyerus/flagd/issues/337)) ([ffb8dc1](https://github.com/skyerus/flagd/commit/ffb8dc167974e7edfa03a2046ca73a85a5e943d6))
* attach image sbom to release artefacts ([#407](https://github.com/skyerus/flagd/issues/407)) ([fb4ee50](https://github.com/skyerus/flagd/commit/fb4ee502217e2262849df09258f3a0ffa7edec13))
* auto update config docs ([#305](https://github.com/skyerus/flagd/issues/305)) ([7d24060](https://github.com/skyerus/flagd/commit/7d240608cd7ac0ef1fad0dc7ae63682164d451c0))
* buf upgrade to v1 ([#229](https://github.com/skyerus/flagd/issues/229)) ([af6f133](https://github.com/skyerus/flagd/commit/af6f133688b506f52ac6180890e3d785238d7b40))
* bump go version ([#227](https://github.com/skyerus/flagd/issues/227)) ([4b9de42](https://github.com/skyerus/flagd/commit/4b9de42a35facb6b66babab3350dd4e34162a843))
* bundle release assets tag_name ([#376](https://github.com/skyerus/flagd/issues/376)) ([0880651](https://github.com/skyerus/flagd/commit/0880651426324728f7446ddd88ac5523b91b9fc4))
* Configure Renovate ([#203](https://github.com/skyerus/flagd/issues/203)) ([fd3b8a0](https://github.com/skyerus/flagd/commit/fd3b8a0c2431ee436744810b61730c9252ce6efe))
* contributing section for the docs ([#202](https://github.com/skyerus/flagd/issues/202)) ([187f0f9](https://github.com/skyerus/flagd/commit/187f0f906e1c443208acf1f39026f542cbd3da2b))
* **deps:** pin anchore/sbom-action action to 06e1094 ([#354](https://github.com/skyerus/flagd/issues/354)) ([c59f49d](https://github.com/skyerus/flagd/commit/c59f49d4479e9477562908cc0a8e3e8030af7b69))
* **deps:** pin dependencies ([#258](https://github.com/skyerus/flagd/issues/258)) ([9cef96c](https://github.com/skyerus/flagd/commit/9cef96c85426a4f0755c4674c7998320918bb914))
* **deps:** pin dependencies ([#473](https://github.com/skyerus/flagd/issues/473)) ([679e860](https://github.com/skyerus/flagd/commit/679e8600f57ab1e03c493c4a4046bd9d7368efac))
* **deps:** pin softprops/action-gh-release action to de2c0eb ([#364](https://github.com/skyerus/flagd/issues/364)) ([a767627](https://github.com/skyerus/flagd/commit/a767627b44fd8b04d2919d30f4daf2cd87de213d))
* **deps:** update actions/checkout action to v3 ([#277](https://github.com/skyerus/flagd/issues/277)) ([be6496b](https://github.com/skyerus/flagd/commit/be6496b95dfe16d619e344951d2a50d36d5fc1e0))
* **deps:** update actions/configure-pages action to v3 ([#294](https://github.com/skyerus/flagd/issues/294)) ([7410191](https://github.com/skyerus/flagd/commit/7410191f491acf3042b8b0c65e9606e92069d6aa))
* **deps:** update actions/configure-pages digest to 37538b2 ([#363](https://github.com/skyerus/flagd/issues/363)) ([f146171](https://github.com/skyerus/flagd/commit/f1461711a11797ef68b0169292954f2aabc944b1))
* **deps:** update actions/configure-pages digest to 3c444c0 ([#349](https://github.com/skyerus/flagd/issues/349)) ([cb84cd8](https://github.com/skyerus/flagd/commit/cb84cd86a9d3dbb7e280e9e1d5409b8b5b807969))
* **deps:** update actions/configure-pages digest to fc89b04 ([#417](https://github.com/skyerus/flagd/issues/417)) ([04014e7](https://github.com/skyerus/flagd/commit/04014e7cb37e43f5ed3726dfd31da96202abc043))
* **deps:** update actions/deploy-pages digest to 0243b6c ([#390](https://github.com/skyerus/flagd/issues/390)) ([9613ea2](https://github.com/skyerus/flagd/commit/9613ea266b7f183cbbe3422ff952067002f5cc85))
* **deps:** update actions/setup-go action to v3 ([#278](https://github.com/skyerus/flagd/issues/278)) ([96a71e5](https://github.com/skyerus/flagd/commit/96a71e5265e1a4143441a8fc44f8d6ad3e96b6dd))
* **deps:** update amannn/action-semantic-pull-request action to v5 ([#279](https://github.com/skyerus/flagd/issues/279)) ([4f79bf9](https://github.com/skyerus/flagd/commit/4f79bf909eec6177a403cbc7475b552e3acaa1ed))
* **deps:** update amannn/action-semantic-pull-request digest to b6bca70 ([#441](https://github.com/skyerus/flagd/issues/441)) ([ce0ebe1](https://github.com/skyerus/flagd/commit/ce0ebe13dd992688a3a0464ff401f2c40651da52))
* **deps:** update anchore/sbom-action digest to 07978da ([#392](https://github.com/skyerus/flagd/issues/392)) ([e4572ee](https://github.com/skyerus/flagd/commit/e4572ee5ac1bdb327208fe106baad5e9c32a369a))
* **deps:** update docker/build-push-action action to v3 ([#280](https://github.com/skyerus/flagd/issues/280)) ([a19d63e](https://github.com/skyerus/flagd/commit/a19d63edc46a133556004e619b9168ebea9b1ea2))
* **deps:** update docker/build-push-action action to v4 ([#351](https://github.com/skyerus/flagd/issues/351)) ([a048750](https://github.com/skyerus/flagd/commit/a0487500c0fa9265c3fda9d34f47ab27131754d4))
* **deps:** update docker/build-push-action digest to 1104d47 ([#350](https://github.com/skyerus/flagd/issues/350)) ([02c47f1](https://github.com/skyerus/flagd/commit/02c47f1c6374fe2ef54083ce6f7e6561ea3f3df2))
* **deps:** update docker/build-push-action digest to 37abced ([#323](https://github.com/skyerus/flagd/issues/323)) ([8b9b948](https://github.com/skyerus/flagd/commit/8b9b94881a6808b1d3cce4d4cc57d3f49a2e596b))
* **deps:** update docker/login-action digest to 3da7dc6 ([#254](https://github.com/skyerus/flagd/issues/254)) ([5ddc8f7](https://github.com/skyerus/flagd/commit/5ddc8f79584206ae481083901521ed371dd764b2))
* **deps:** update docker/login-action digest to ec9cdf0 ([#437](https://github.com/skyerus/flagd/issues/437)) ([2650670](https://github.com/skyerus/flagd/commit/2650670d35166e119f9a92613d3aca81523b9faa))
* **deps:** update docker/metadata-action digest to 05d22bf ([#270](https://github.com/skyerus/flagd/issues/270)) ([d43ee80](https://github.com/skyerus/flagd/commit/d43ee80a1f4bd942800e1d47913da9699ef386ef))
* **deps:** update docker/metadata-action digest to 3343011 ([#438](https://github.com/skyerus/flagd/issues/438)) ([e7ebf32](https://github.com/skyerus/flagd/commit/e7ebf32caf0eae7449e673da0c10998f97ebf781))
* **deps:** update docker/metadata-action digest to 507c2f2 ([#287](https://github.com/skyerus/flagd/issues/287)) ([f5b0ece](https://github.com/skyerus/flagd/commit/f5b0eceffe6ff712b57e5bc2abbf584b05771107))
* **deps:** update docker/metadata-action digest to 6c3ca5d ([#289](https://github.com/skyerus/flagd/issues/289)) ([d042e1f](https://github.com/skyerus/flagd/commit/d042e1f4a3236a908b90b796df5a7095c58a98b2))
* **deps:** update docker/metadata-action digest to 766400c ([#355](https://github.com/skyerus/flagd/issues/355)) ([aaa19e3](https://github.com/skyerus/flagd/commit/aaa19e328a326a3a26f8d659de9fc6b340a13e81))
* **deps:** update docker/metadata-action digest to 9e6ae02 ([#255](https://github.com/skyerus/flagd/issues/255)) ([ceb1a84](https://github.com/skyerus/flagd/commit/ceb1a84820200e943bd97c372cdbe85e8f012504))
* **deps:** update github/codeql-action digest to 17573ee ([#403](https://github.com/skyerus/flagd/issues/403)) ([9d7dbd1](https://github.com/skyerus/flagd/commit/9d7dbd1b0d81cfa67c99aa55f483517f3d6987ca))
* **deps:** update github/codeql-action digest to 32dc499 ([#439](https://github.com/skyerus/flagd/issues/439)) ([f91d91b](https://github.com/skyerus/flagd/commit/f91d91bf020d330f96572c5ee11a210c0c7f4311))
* **deps:** update github/codeql-action digest to 39d8d7e ([#382](https://github.com/skyerus/flagd/issues/382)) ([446e269](https://github.com/skyerus/flagd/commit/446e2697f4411ea5e03f9ddd04182051e08f31d9))
* **deps:** update github/codeql-action digest to 3ebbd71 ([#333](https://github.com/skyerus/flagd/issues/333)) ([d64e245](https://github.com/skyerus/flagd/commit/d64e2450cfdfc313b4cad169e8336f7d0c29cb03))
* **deps:** update github/codeql-action digest to 436dbd9 ([#330](https://github.com/skyerus/flagd/issues/330)) ([3b948c1](https://github.com/skyerus/flagd/commit/3b948c18b9edf7334cdc32d85bb4e51ff9b1e9c8))
* **deps:** update github/codeql-action digest to 515828d ([#281](https://github.com/skyerus/flagd/issues/281)) ([aaa066b](https://github.com/skyerus/flagd/commit/aaa066b04c0f4f4695e7c8e870450e4fe5297e9a))
* **deps:** update github/codeql-action digest to 8775e86 ([#395](https://github.com/skyerus/flagd/issues/395)) ([49dbe80](https://github.com/skyerus/flagd/commit/49dbe80b6f8c794e22c3101dc0f71a6bc4e9bc7d))
* **deps:** update github/codeql-action digest to a34ca99 ([#292](https://github.com/skyerus/flagd/issues/292)) ([9db4186](https://github.com/skyerus/flagd/commit/9db4186482483d790169a887093f0419a01f6fdc))
* **deps:** update golang docker tag to v1.19 ([#266](https://github.com/skyerus/flagd/issues/266)) ([36e0f28](https://github.com/skyerus/flagd/commit/36e0f281437105ed6f6af80506a094cb6b2c501e))
* **deps:** update golang docker tag to v1.20 ([#366](https://github.com/skyerus/flagd/issues/366)) ([eaa62b3](https://github.com/skyerus/flagd/commit/eaa62b371adb1a0d70873e74409ed3e7505f8305))
* **deps:** update google-github-actions/release-please-action digest to 514bc6c ([#393](https://github.com/skyerus/flagd/issues/393)) ([73905cc](https://github.com/skyerus/flagd/commit/73905ccba9853e5fe202e8179676a7001d0cca2f))
* **deps:** update google-github-actions/release-please-action digest to d3c71f9 ([#406](https://github.com/skyerus/flagd/issues/406)) ([6e1ffb2](https://github.com/skyerus/flagd/commit/6e1ffb27fea5e91014a6991b2afca9a59f89117f))
* **deps:** update google-github-actions/release-please-action digest to e0b9d18 ([#474](https://github.com/skyerus/flagd/issues/474)) ([5b85b2a](https://github.com/skyerus/flagd/commit/5b85b2a611d9199e39735f101ed7e560257ce2e4))
* **deps:** update goreleaser/goreleaser-action digest to f82d6c1 ([#347](https://github.com/skyerus/flagd/issues/347)) ([93def53](https://github.com/skyerus/flagd/commit/93def53adf958fcb3b570876fa8189ac0cd5f378))
* **deps:** update sigstore/cosign-installer digest to bd2d118 ([#471](https://github.com/skyerus/flagd/issues/471)) ([ee90f48](https://github.com/skyerus/flagd/commit/ee90f48317ec600f09534306503dc752254a1d09))
* disable caching tests in CI ([#442](https://github.com/skyerus/flagd/issues/442)) ([28a35f6](https://github.com/skyerus/flagd/commit/28a35f62d618539362ae83a48f11af08ca2ae245))
* exclude component in git tag ([#190](https://github.com/skyerus/flagd/issues/190)) ([84e32b7](https://github.com/skyerus/flagd/commit/84e32b72525b4cfb05fc7df4d82b2983fe3d33c3))
* fix benchmark tests & ensure that a failure of benchmarks is evident in ci ([#295](https://github.com/skyerus/flagd/issues/295)) ([ce14db2](https://github.com/skyerus/flagd/commit/ce14db257e08acb0e3990e706f336ae196493655))
* fix race condition on init read ([#409](https://github.com/skyerus/flagd/issues/409)) ([0c9eb23](https://github.com/skyerus/flagd/commit/0c9eb2322df99b4216d40afd1cb3b8873b0c59ff))
* Image source for image ([#348](https://github.com/skyerus/flagd/issues/348)) ([0391848](https://github.com/skyerus/flagd/commit/0391848317f615253190176b25ad854569a391c5))
* integration test stability ([#432](https://github.com/skyerus/flagd/issues/432)) ([5a6a5d5](https://github.com/skyerus/flagd/commit/5a6a5d5887badd846cffe882c8c22a35b850fa06))
* integration tests ([#312](https://github.com/skyerus/flagd/issues/312)) ([6192ac8](https://github.com/skyerus/flagd/commit/6192ac8820b0f472672ba177b7c5838244b6e277))
* **main:** release 0.1.1 ([#164](https://github.com/skyerus/flagd/issues/164)) ([7559754](https://github.com/skyerus/flagd/commit/7559754bf865c816937af0427af4068eb4e03c96))
* **main:** release 0.2.0 ([#172](https://github.com/skyerus/flagd/issues/172)) ([c624173](https://github.com/skyerus/flagd/commit/c624173bf39c53061435ff5614fc67235bdd3f38))
* **main:** release 0.2.1 ([#175](https://github.com/skyerus/flagd/issues/175)) ([087975d](https://github.com/skyerus/flagd/commit/087975d3d2542d1ff887a25d9e74101fe82ce65c))
* **main:** release 0.2.2 ([#183](https://github.com/skyerus/flagd/issues/183)) ([3c78cff](https://github.com/skyerus/flagd/commit/3c78cff2186daeaf31801faa28e154be75df7487))
* **main:** release 0.2.3 ([#185](https://github.com/skyerus/flagd/issues/185)) ([ccb629e](https://github.com/skyerus/flagd/commit/ccb629ebd0b4799983558a411cb7f9a4c321748f))
* **main:** release 0.2.4 ([#194](https://github.com/skyerus/flagd/issues/194)) ([24c8984](https://github.com/skyerus/flagd/commit/24c8984bf76e75a6e112eeffa3809eb0e1ee9ce3))
* **main:** release 0.2.5 ([#199](https://github.com/skyerus/flagd/issues/199)) ([98a7397](https://github.com/skyerus/flagd/commit/98a7397c3f9d011b9aaae26ef3a2e85c27a82e51))
* **main:** release 0.2.7 ([#215](https://github.com/skyerus/flagd/issues/215)) ([df7c6ee](https://github.com/skyerus/flagd/commit/df7c6ee15ab30325b893cdce95e0735d486ebf74))
* **main:** release 0.3.0 ([#230](https://github.com/skyerus/flagd/issues/230)) ([2ad206e](https://github.com/skyerus/flagd/commit/2ad206e635cc5b68ae42287b740ddfc86b0194ff))
* **main:** release 0.3.1 ([#253](https://github.com/skyerus/flagd/issues/253)) ([7a1b1ef](https://github.com/skyerus/flagd/commit/7a1b1ef9b5c57b201612a49d3fa0d7405d421dc4))
* **main:** release 0.3.2 ([#284](https://github.com/skyerus/flagd/issues/284)) ([071b180](https://github.com/skyerus/flagd/commit/071b18091f88f96b846891223bce551e98f4310c))
* **main:** release 0.3.3 ([#332](https://github.com/skyerus/flagd/issues/332)) ([745e565](https://github.com/skyerus/flagd/commit/745e5652a3f78169aa665ed64568001383659f98))
* **main:** release 0.3.4 ([#341](https://github.com/skyerus/flagd/issues/341)) ([311cb92](https://github.com/skyerus/flagd/commit/311cb92fc205fcaf0cab4773ea7f8c53bf63cc86))
* **main:** release 0.3.5 ([#343](https://github.com/skyerus/flagd/issues/343)) ([e77950b](https://github.com/skyerus/flagd/commit/e77950bf6e4c337671de745a4a500cb464649079))
* **main:** release 0.3.6 ([#378](https://github.com/skyerus/flagd/issues/378)) ([79ac2f6](https://github.com/skyerus/flagd/commit/79ac2f61b07170af6ae3509986b8e4b600a8681b))
* **main:** release 0.3.7 ([#381](https://github.com/skyerus/flagd/issues/381)) ([b71729a](https://github.com/skyerus/flagd/commit/b71729ae8b06998a9ff9b68c628ad16727bb1033))
* **main:** release 0.4.0 ([#412](https://github.com/skyerus/flagd/issues/412)) ([408bb7c](https://github.com/skyerus/flagd/commit/408bb7c6420678706854932df3144253ada89f82))
* **main:** release 0.4.1 ([#462](https://github.com/skyerus/flagd/issues/462)) ([4e904a2](https://github.com/skyerus/flagd/commit/4e904a2d77c8c2fe05d7a42ecb306ed12d9884e6))
* refactoring and improve coverage for K8s Sync ([#466](https://github.com/skyerus/flagd/issues/466)) ([6dc441e](https://github.com/skyerus/flagd/commit/6dc441e2f2418c1fd3a5a58dbb99f848ccbd8735))
* release v0.2.6 ([93cfb78](https://github.com/skyerus/flagd/commit/93cfb78d024b436fa7fb17fd41f74d1508bf8b64))
* release v0.2.7 ([4a9f6df](https://github.com/skyerus/flagd/commit/4a9f6df4e472229ff805e9d5d3aa581c7c9c0667))
* remove test log ([#384](https://github.com/skyerus/flagd/issues/384)) ([ff518a3](https://github.com/skyerus/flagd/commit/ff518a39302445b7ddb3778e4bc6d32d68a8caec))
* remove unused code ([#235](https://github.com/skyerus/flagd/issues/235)) ([0749972](https://github.com/skyerus/flagd/commit/0749972c6e7b856ed78b1947548616cccdb7fbaa))
* renovate, auto-merge gh actions ([978c2f0](https://github.com/skyerus/flagd/commit/978c2f0f36de973cccac7b5b3448c6dd68c12050))
* reorder release note sections ([df7bfce](https://github.com/skyerus/flagd/commit/df7bfce85ec7d6abaa987f87341c5af380904b51))
* update architecture doc image path ([c2a9ac9](https://github.com/skyerus/flagd/commit/c2a9ac9c72e775b0da2e055c850d5288b0d2708c))
* update CODEOWNERS ([#228](https://github.com/skyerus/flagd/issues/228)) ([3ecfc65](https://github.com/skyerus/flagd/commit/3ecfc654f92cff7ce6ba7f802d5196202f88c653))
* update link to ofep ([cc860bf](https://github.com/skyerus/flagd/commit/cc860bfdc1bedae69349634fd4d3869120f78a0f))
* update readme to include a features section ([#195](https://github.com/skyerus/flagd/issues/195)) ([a41d84f](https://github.com/skyerus/flagd/commit/a41d84f5bb850fa404eb5c85ae89c7c746a90af8))
* use -short flag in benchmark tests ([#431](https://github.com/skyerus/flagd/issues/431)) ([e68a6aa](https://github.com/skyerus/flagd/commit/e68a6aadb3dac46676299ab94a34a0bcc39a67af))

## [0.4.1](https://github.com/open-feature/flagd/compare/v0.4.0...v0.4.1) (2023-03-07)


### 🔄 Refactoring

* remove unused struct field ([#458](https://github.com/open-feature/flagd/issues/458)) ([a04c0b8](https://github.com/open-feature/flagd/commit/a04c0b837dbe9e28d1e01e43ea9e378a6c0f316a))


### 🧹 Chore

* **deps:** update sigstore/cosign-installer digest to bd2d118 ([#471](https://github.com/open-feature/flagd/issues/471)) ([ee90f48](https://github.com/open-feature/flagd/commit/ee90f48317ec600f09534306503dc752254a1d09))


### 🐛 Bug Fixes

* **deps:** update module github.com/open-feature/go-sdk-contrib/providers/flagd to v0.1.10 ([#459](https://github.com/open-feature/flagd/issues/459)) ([cbdf9b0](https://github.com/open-feature/flagd/commit/cbdf9b07c30239d7d04ef770cf4461fb33422fe9))
* **deps:** update module golang.org/x/net to v0.8.0 ([#468](https://github.com/open-feature/flagd/issues/468)) ([10d5f2c](https://github.com/open-feature/flagd/commit/10d5f2c55081a25daa1f0e0fa81f96f0fffbbc7b))
* fix broken image signing ([#461](https://github.com/open-feature/flagd/issues/461)) ([05bb51c](https://github.com/open-feature/flagd/commit/05bb51c7ab30f6e976b87f54ca889e978f834211))
* fixing image delimeter  ([#463](https://github.com/open-feature/flagd/issues/463)) ([b4ee495](https://github.com/open-feature/flagd/commit/b4ee495dc8e00b032518ea42d272a36b3b662e95))
* security issues ([#464](https://github.com/open-feature/flagd/issues/464)) ([7f1e759](https://github.com/open-feature/flagd/commit/7f1e759a87a9af63e9384005c959a3f500cc474c))
* set readiness once only ([#465](https://github.com/open-feature/flagd/issues/465)) ([41a888d](https://github.com/open-feature/flagd/commit/41a888d6b60c030b913280c2a1eeff8b25e8aada))

## [0.4.0](https://github.com/open-feature/flagd/compare/v0.3.7...v0.4.0) (2023-03-02)


### ⚠ BREAKING CHANGES

* Use OTel to export metrics (metric name changes) ([#419](https://github.com/open-feature/flagd/issues/419))

### 🧹 Chore

* add additional sections to the release notes ([#449](https://github.com/open-feature/flagd/issues/449)) ([798f71a](https://github.com/open-feature/flagd/commit/798f71a92d2e2f450a53cda93b44217cbb2ad7fd))
* attach image sbom to release artefacts ([#407](https://github.com/open-feature/flagd/issues/407)) ([fb4ee50](https://github.com/open-feature/flagd/commit/fb4ee502217e2262849df09258f3a0ffa7edec13))
* **deps:** update actions/configure-pages digest to fc89b04 ([#417](https://github.com/open-feature/flagd/issues/417)) ([04014e7](https://github.com/open-feature/flagd/commit/04014e7cb37e43f5ed3726dfd31da96202abc043))
* **deps:** update amannn/action-semantic-pull-request digest to b6bca70 ([#441](https://github.com/open-feature/flagd/issues/441)) ([ce0ebe1](https://github.com/open-feature/flagd/commit/ce0ebe13dd992688a3a0464ff401f2c40651da52))
* **deps:** update docker/login-action digest to ec9cdf0 ([#437](https://github.com/open-feature/flagd/issues/437)) ([2650670](https://github.com/open-feature/flagd/commit/2650670d35166e119f9a92613d3aca81523b9faa))
* **deps:** update docker/metadata-action digest to 3343011 ([#438](https://github.com/open-feature/flagd/issues/438)) ([e7ebf32](https://github.com/open-feature/flagd/commit/e7ebf32caf0eae7449e673da0c10998f97ebf781))
* **deps:** update github/codeql-action digest to 32dc499 ([#439](https://github.com/open-feature/flagd/issues/439)) ([f91d91b](https://github.com/open-feature/flagd/commit/f91d91bf020d330f96572c5ee11a210c0c7f4311))
* **deps:** update google-github-actions/release-please-action digest to d3c71f9 ([#406](https://github.com/open-feature/flagd/issues/406)) ([6e1ffb2](https://github.com/open-feature/flagd/commit/6e1ffb27fea5e91014a6991b2afca9a59f89117f))
* disable caching tests in CI ([#442](https://github.com/open-feature/flagd/issues/442)) ([28a35f6](https://github.com/open-feature/flagd/commit/28a35f62d618539362ae83a48f11af08ca2ae245))
* fix race condition on init read ([#409](https://github.com/open-feature/flagd/issues/409)) ([0c9eb23](https://github.com/open-feature/flagd/commit/0c9eb2322df99b4216d40afd1cb3b8873b0c59ff))
* integration test stability ([#432](https://github.com/open-feature/flagd/issues/432)) ([5a6a5d5](https://github.com/open-feature/flagd/commit/5a6a5d5887badd846cffe882c8c22a35b850fa06))
* integration tests ([#312](https://github.com/open-feature/flagd/issues/312)) ([6192ac8](https://github.com/open-feature/flagd/commit/6192ac8820b0f472672ba177b7c5838244b6e277))
* reorder release note sections ([df7bfce](https://github.com/open-feature/flagd/commit/df7bfce85ec7d6abaa987f87341c5af380904b51))
* use -short flag in benchmark tests ([#431](https://github.com/open-feature/flagd/issues/431)) ([e68a6aa](https://github.com/open-feature/flagd/commit/e68a6aadb3dac46676299ab94a34a0bcc39a67af))


### 🐛 Bug Fixes

* **deps:** update kubernetes packages to v0.26.2 ([#450](https://github.com/open-feature/flagd/issues/450)) ([2885227](https://github.com/open-feature/flagd/commit/28852270f34ff81c072337b29aa17f4b6634e9cc))
* **deps:** update module github.com/bufbuild/connect-go to v1.5.2 ([#416](https://github.com/open-feature/flagd/issues/416)) ([feb7f04](https://github.com/open-feature/flagd/commit/feb7f047365263758a63d8dffea936f621a4966d))
* **deps:** update module github.com/open-feature/go-sdk-contrib/providers/flagd to v0.1.9 ([#427](https://github.com/open-feature/flagd/issues/427)) ([42d2705](https://github.com/open-feature/flagd/commit/42d270558bf9badcff9a9b352fda35491c45aebe))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.29 ([#429](https://github.com/open-feature/flagd/issues/429)) ([b7fae81](https://github.com/open-feature/flagd/commit/b7fae81b89b3a1a0793a688c32569c4284633c6a))
* **deps:** update module github.com/stretchr/testify to v1.8.2 ([#440](https://github.com/open-feature/flagd/issues/440)) ([ab3e674](https://github.com/open-feature/flagd/commit/ab3e6748abc7843c022afeaf7cb11193cdcf59c5))
* **deps:** update module golang.org/x/net to v0.7.0 ([#410](https://github.com/open-feature/flagd/issues/410)) ([c6133b6](https://github.com/open-feature/flagd/commit/c6133b6af61f3d73ae73d318a1a9f44db2540540))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.5 ([#454](https://github.com/open-feature/flagd/issues/454)) ([f907f11](https://github.com/open-feature/flagd/commit/f907f114f23fa2efa2637e254e829e4d53a90b51))
* remove non-error error log from parseFractionalEvaluationData ([#446](https://github.com/open-feature/flagd/issues/446)) ([34aca79](https://github.com/open-feature/flagd/commit/34aca79e6ec9876a6cced0fe49e1ceea34d83696))


### ✨ New Features

* add debug logging for merge behaviour  ([#456](https://github.com/open-feature/flagd/issues/456)) ([dc71e84](https://github.com/open-feature/flagd/commit/dc71e84f0704690b528e7f1c2b56cb4898374fbf))
* add Health and Readiness probes ([#418](https://github.com/open-feature/flagd/issues/418)) ([7f2358c](https://github.com/open-feature/flagd/commit/7f2358ce207527c890f4a2f46ce4b9e8bf697095))
* Add version to startup message ([#430](https://github.com/open-feature/flagd/issues/430)) ([8daf613](https://github.com/open-feature/flagd/commit/8daf613e7e4f4492df0c06e2ef464f4337cadaca))
* introduce flag merge behaviour ([#414](https://github.com/open-feature/flagd/issues/414)) ([524f65e](https://github.com/open-feature/flagd/commit/524f65ea7215466bb4ac24a8d0d5953dd1cfe9a0))
* introduce grpc sync for flagd ([#297](https://github.com/open-feature/flagd/issues/297)) ([33413f2](https://github.com/open-feature/flagd/commit/33413f25882a3f1cf4953da0f18e746bfb69faf4))
* refactor and improve K8s sync provider ([#443](https://github.com/open-feature/flagd/issues/443)) ([4c03bfc](https://github.com/open-feature/flagd/commit/4c03bfc812e7ceabcac0979290bd74d9efc9da15))
* Use OTel to export metrics (metric name changes) ([#419](https://github.com/open-feature/flagd/issues/419)) ([eb3982a](https://github.com/open-feature/flagd/commit/eb3982a1cb72d664022b5cb126b533cf61497001))


### 📚 Documentation

* add .net flagd provider ([73d7840](https://github.com/open-feature/flagd/commit/73d7840c9fdef9c62371c677e02c0d9773c85f95))
* configuration merge docs ([#455](https://github.com/open-feature/flagd/issues/455)) ([6cb66b1](https://github.com/open-feature/flagd/commit/6cb66b14d01b6ee1c270bbdd3e30d4016757eae5))
* documentation for creating a provider ([#413](https://github.com/open-feature/flagd/issues/413)) ([d0c099d](https://github.com/open-feature/flagd/commit/d0c099d9aba3ed4d760a1858381f5e29b6d49a9c))
* updated filepaths for schema store regex ([#344](https://github.com/open-feature/flagd/issues/344)) ([2d0e9d9](https://github.com/open-feature/flagd/commit/2d0e9d956fbc99f2775821cfecdceb2b016d2b78))

## [0.3.7](https://github.com/open-feature/flagd/compare/v0.3.6...v0.3.7) (2023-02-13)


### Bug Fixes

* **deps:** update module golang.org/x/net to v0.6.0 ([#396](https://github.com/open-feature/flagd/issues/396)) ([beb7564](https://github.com/open-feature/flagd/commit/beb756470b1e1d5ef0670b8322b6ed9cb44efa24))
* **deps:** update module google.golang.org/grpc to v1.53.0 ([#388](https://github.com/open-feature/flagd/issues/388)) ([174cd7c](https://github.com/open-feature/flagd/commit/174cd7c70fa5ae2573db2c5972b75786633e2f41))
* error handling of Serve/ServeTLS funcs ([#397](https://github.com/open-feature/flagd/issues/397)) ([8923bf2](https://github.com/open-feature/flagd/commit/8923bf2d407e18b65c188aef9bf7370fc74c3be2))
* fix race in http sync test ([#401](https://github.com/open-feature/flagd/issues/401)) ([1d0c8e1](https://github.com/open-feature/flagd/commit/1d0c8e168b73f7fbd4b27ece733041bbe08261c0))
* sbom artefact name ([#380](https://github.com/open-feature/flagd/issues/380)) ([3daef26](https://github.com/open-feature/flagd/commit/3daef263c43ed63776d604d27f7ae6b993fff143)), closes [#379](https://github.com/open-feature/flagd/issues/379)

## [0.3.6](https://github.com/open-feature/flagd/compare/v0.3.5...v0.3.6) (2023-02-06)


### Bug Fixes

* set ResolveObject reason ([#375](https://github.com/open-feature/flagd/issues/375)) ([dcf199d](https://github.com/open-feature/flagd/commit/dcf199dab9d11b86454028869a54d77a474fc4a6))

## [0.3.5](https://github.com/open-feature/flagd/compare/v0.3.4...v0.3.5) (2023-02-06)


### Features

* flagd image signing ([#338](https://github.com/open-feature/flagd/issues/338)) ([eca6a60](https://github.com/open-feature/flagd/commit/eca6a60967999a303ceef5465f1acc35c83afd6d))
* update in logging to console and Unify case usage, seperators and punctuation for logging ([#322](https://github.com/open-feature/flagd/issues/322)) ([0bdcfd2](https://github.com/open-feature/flagd/commit/0bdcfd2fecc03b15be9fc4b0489431b8fa86aed8))


### Bug Fixes

* **deps:** update module github.com/bufbuild/connect-go to v1.5.1 ([#365](https://github.com/open-feature/flagd/issues/365)) ([e25f452](https://github.com/open-feature/flagd/commit/e25f452906e034e339309270cc8db6dcd58e9973))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.28 ([#342](https://github.com/open-feature/flagd/issues/342)) ([e6df80f](https://github.com/open-feature/flagd/commit/e6df80fd25d3da342e72d2ca0e923d9bf3d3f797))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.2 ([#336](https://github.com/open-feature/flagd/issues/336)) ([836d3cf](https://github.com/open-feature/flagd/commit/836d3cf3c06570d59929c3464e3c8e11c9b5a2fa))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.3 ([#372](https://github.com/open-feature/flagd/issues/372)) ([330ac91](https://github.com/open-feature/flagd/commit/330ac91e375124826b2a7a1a22d0daa18368ab99))
* **deps:** update module sigs.k8s.io/controller-runtime to v0.14.4 ([#374](https://github.com/open-feature/flagd/issues/374)) ([d90e561](https://github.com/open-feature/flagd/commit/d90e561bfc5b798d13d4ba8f30f523b1053f3748))
* fix unbuffered channel blocking goroutine  ([#358](https://github.com/open-feature/flagd/issues/358)) ([4f1905a](https://github.com/open-feature/flagd/commit/4f1905a9ac6d62b5edb297fba904aac8680c89cf))
* introduced RWMutex to flag state to prevent concurrent r/w of map ([#370](https://github.com/open-feature/flagd/issues/370)) ([93e356b](https://github.com/open-feature/flagd/commit/93e356b4ab0b65c71659bd52d73f618edffc96f5))
* use event.Has func for file change notification handling (increased stability across OS) ([#361](https://github.com/open-feature/flagd/issues/361)) ([09f74b9](https://github.com/open-feature/flagd/commit/09f74b9c5d15622c98da08558cbcd63fe9422754))

## [0.3.4](https://github.com/open-feature/flagd/compare/v0.3.3...v0.3.4) (2023-01-28)


### Bug Fixes

* **deps:** update goreleaser/goreleaser-action action to v4 ([#340](https://github.com/open-feature/flagd/issues/340)) ([b9fcd5c](https://github.com/open-feature/flagd/commit/b9fcd5caa67a61b447c437b651471b4603b2b272))

## [0.3.3](https://github.com/open-feature/flagd/compare/v0.3.2...v0.3.3) (2023-01-28)


### Bug Fixes

* **deps:** update module github.com/bufbuild/connect-go to v1.5.0 ([#326](https://github.com/open-feature/flagd/issues/326)) ([7f332e5](https://github.com/open-feature/flagd/commit/7f332e50ecb1cea19108d1fa2fd79da3d5864bf9))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.26 ([#331](https://github.com/open-feature/flagd/issues/331)) ([be67e5f](https://github.com/open-feature/flagd/commit/be67e5f5bc1fb7351a04ffc4180447a27d57d32a))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.27 ([#335](https://github.com/open-feature/flagd/issues/335)) ([824cf1a](https://github.com/open-feature/flagd/commit/824cf1ab0f2e18826207af16d5328b817c755c8e))
* send datasync on remove fs events ([#339](https://github.com/open-feature/flagd/issues/339)) ([4c9aaac](https://github.com/open-feature/flagd/commit/4c9aaaca77b1c8b16f59434aeb37407fced47ecf))

## [0.3.2](https://github.com/open-feature/flagd/compare/v0.3.1...v0.3.2) (2023-01-26)


### Bug Fixes

* deprecation warning fix ([#317](https://github.com/open-feature/flagd/issues/317)) ([a2630db](https://github.com/open-feature/flagd/commit/a2630dbba151f35cc985d38de9cf25bfee2b76c8))
* **deps:** update kubernetes packages to v0.26.1 ([#267](https://github.com/open-feature/flagd/issues/267)) ([26825f2](https://github.com/open-feature/flagd/commit/26825f288c56df638fd160caa93f926a8c136108))
* **deps:** update module github.com/diegoholiveira/jsonlogic/v3 to v3.2.7 ([#283](https://github.com/open-feature/flagd/issues/283)) ([2ab5a00](https://github.com/open-feature/flagd/commit/2ab5a00fa6f19c7e0fe1a4e36649fae2633ac269))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.24 ([#290](https://github.com/open-feature/flagd/issues/290)) ([38d3eba](https://github.com/open-feature/flagd/commit/38d3ebaffcb1f36a38003273c62c6317f0ee75a3))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.25 ([#324](https://github.com/open-feature/flagd/issues/324)) ([ed1d3aa](https://github.com/open-feature/flagd/commit/ed1d3aaba4ca179a89757a6b1c3f328826e787fc))
* **deps:** update module github.com/spf13/viper to v1.15.0 ([#296](https://github.com/open-feature/flagd/issues/296)) ([d43220b](https://github.com/open-feature/flagd/commit/d43220b2be58e4bce05050c5d1b36788289ae7cc))
* **deps:** update module google.golang.org/grpc to v1.52.1 ([#314](https://github.com/open-feature/flagd/issues/314)) ([ad25388](https://github.com/open-feature/flagd/commit/ad25388461816100e19bda44a8e0077770ea0ee4))
* **deps:** update module google.golang.org/grpc to v1.52.3 ([#325](https://github.com/open-feature/flagd/issues/325)) ([8013ea5](https://github.com/open-feature/flagd/commit/8013ea5c6fa311b337c7ec1b1e8e756080808948))
* Update flagd systemd config to use URI ([#315](https://github.com/open-feature/flagd/issues/315)) ([93a04b4](https://github.com/open-feature/flagd/commit/93a04b46133e9220ec6f23d833c11f195e05c13e))
* update outdated doc link in deprecation warning ([#316](https://github.com/open-feature/flagd/issues/316)) ([19695d2](https://github.com/open-feature/flagd/commit/19695d2715129d6718ca0617b6ec6922ffb79c9b))

## [0.3.1](https://github.com/open-feature/flagd/compare/v0.3.0...v0.3.1) (2023-01-12)


### Features

* file extension detection ([#257](https://github.com/open-feature/flagd/issues/257)) ([ca22541](https://github.com/open-feature/flagd/commit/ca2254117adc163b94d662b3d1fbfd868f788fcb))
* ResolveAll endpoint for bulk evaluation ([#239](https://github.com/open-feature/flagd/issues/239)) ([6437c43](https://github.com/open-feature/flagd/commit/6437c43022b5c94d2fb835a406d85a4e836f2fcf))


### Bug Fixes

* **deps:** update module github.com/bufbuild/connect-go to v1.4.1 ([#268](https://github.com/open-feature/flagd/issues/268)) ([712d7dd](https://github.com/open-feature/flagd/commit/712d7dd4a34980bf9eddad99d926cbdd5d69d624))
* **deps:** update module github.com/mattn/go-colorable to v0.1.13 ([#260](https://github.com/open-feature/flagd/issues/260)) ([5b11504](https://github.com/open-feature/flagd/commit/5b11504cdce50c137540cc79d2db94e70a21338b))
* **deps:** update module github.com/open-feature/open-feature-operator to v0.2.23 ([#261](https://github.com/open-feature/flagd/issues/261)) ([a1dd3b9](https://github.com/open-feature/flagd/commit/a1dd3b9005374b5527f12b8e138250cacddc71af))
* **deps:** update module github.com/rs/cors to v1.8.3 ([#264](https://github.com/open-feature/flagd/issues/264)) ([0e6f2f3](https://github.com/open-feature/flagd/commit/0e6f2f3e5a77dae7d491eaf1094a65e692bebe5d))
* **deps:** update module github.com/stretchr/testify to v1.8.1 ([#265](https://github.com/open-feature/flagd/issues/265)) ([2ec61c6](https://github.com/open-feature/flagd/commit/2ec61c6bc61c266451b496ff18c3dd9a74173233))
* improve invalid sync URI errror msg ([#252](https://github.com/open-feature/flagd/issues/252)) ([5939870](https://github.com/open-feature/flagd/commit/5939870b8994dbca585c53dd022485090aab2406))
* replace character slice with regex replace ([#250](https://github.com/open-feature/flagd/issues/250)) ([c92d101](https://github.com/open-feature/flagd/commit/c92d1012b0de6af694c3af2fede28053e2572b04))

## [0.3.0](https://github.com/open-feature/flagd/compare/v0.2.7...v0.3.0) (2023-01-06)


### ⚠ BREAKING CHANGES

* consolidated configuration change events into one event ([#241](https://github.com/open-feature/flagd/issues/241))

### Features

* consolidated configuration change events into one event ([#241](https://github.com/open-feature/flagd/issues/241)) ([f9684b8](https://github.com/open-feature/flagd/commit/f9684b858dfef40576e0031654b421a37e8bb1d6))
* support yaml evaluator ([#206](https://github.com/open-feature/flagd/issues/206)) ([2dbace5](https://github.com/open-feature/flagd/commit/2dbace5b6bb8e187a7d44a3d3ec14190c63b3ae0))


### Bug Fixes

* changed eventing configuration mutex to rwmutex and added missing lock ([#220](https://github.com/open-feature/flagd/issues/220)) ([5bbef9e](https://github.com/open-feature/flagd/commit/5bbef9ea4b1960686e58298c2c2e192ca99f072f))
* omitempty targeting field in Flag structure ([#247](https://github.com/open-feature/flagd/issues/247)) ([3f406b5](https://github.com/open-feature/flagd/commit/3f406b53bda8b5beb8b0929da3802a0368c13151))

## [0.2.7](https://github.com/open-feature/flagd/compare/v0.2.5...v0.2.7) (2022-12-02)


### ⚠ BREAKING CHANGES

* start command flag refactor ([#222](https://github.com/open-feature/flagd/issues/222))

### Features

* enable request logging via the --debug flag ([#226](https://github.com/open-feature/flagd/issues/226)) ([11954b5](https://github.com/open-feature/flagd/commit/11954b521cc6197d0dc04b163e66e38d4c288047))
* Resurrected the STATIC flag reason. Documented the caching strategy. ([#224](https://github.com/open-feature/flagd/issues/224)) ([5830592](https://github.com/open-feature/flagd/commit/5830592053c55dc9e55c16854e40c3fc8345d6d1))
* snap ([#211](https://github.com/open-feature/flagd/issues/211)) ([c619844](https://github.com/open-feature/flagd/commit/c61984448d5cdadec62b5cf6f7e24fc5f75a3738))
* start command flag refactor ([#222](https://github.com/open-feature/flagd/issues/222)) ([14474cc](https://github.com/open-feature/flagd/commit/14474ccf65b9b92213e8c792e94c458022484df4))


### Miscellaneous Chores

* release v0.2.6 ([93cfb78](https://github.com/open-feature/flagd/commit/93cfb78d024b436fa7fb17fd41f74d1508bf8b64))
* release v0.2.7 ([4a9f6df](https://github.com/open-feature/flagd/commit/4a9f6df4e472229ff805e9d5d3aa581c7c9c0667))

## [0.2.5](https://github.com/open-feature/flagd/compare/v0.2.4...v0.2.5) (2022-10-20)


### Bug Fixes

* CVE-2022-32149 ([#198](https://github.com/open-feature/flagd/issues/198)) ([11a7b34](https://github.com/open-feature/flagd/commit/11a7b3472ab2bc39bce7c40037e8f83736065163))

## [0.2.4](https://github.com/open-feature/flagd/compare/v0.2.3...v0.2.4) (2022-10-14)


### Bug Fixes

* ApiVersion check fix ([#193](https://github.com/open-feature/flagd/issues/193)) ([3a524d6](https://github.com/open-feature/flagd/commit/3a524d646187355bb224100f436c7b5f35abea3e))

## [0.2.3](https://github.com/open-feature/flagd/compare/v0.2.2...v0.2.3) (2022-10-13)


### Features

* Eventing ([#187](https://github.com/open-feature/flagd/issues/187)) ([3f7fcd2](https://github.com/open-feature/flagd/commit/3f7fcd2f57318fad4e0bf501f202af990d3c5a79))
* fixing informer issues ([#191](https://github.com/open-feature/flagd/issues/191)) ([837b0c6](https://github.com/open-feature/flagd/commit/837b0c673e7e7d4799f100291ca520d22944f22a))
* only fire modify event when FeatureFlagConfiguration Generation field has changed ([#167](https://github.com/open-feature/flagd/issues/167)) ([e2fc7ee](https://github.com/open-feature/flagd/commit/e2fc7ee2570a119923bf95b40b2046dfa4705f20))

## [0.2.2](https://github.com/open-feature/flagd/compare/v0.2.1...v0.2.2) (2022-10-03)


### Bug Fixes

* updated merge functionality ([#182](https://github.com/open-feature/flagd/issues/182)) ([94d7697](https://github.com/open-feature/flagd/commit/94d7697d08a07cede4a548ef998792d00f8954a0))

## [0.2.1](https://github.com/open-feature/flagd/compare/v0.2.0...v0.2.1) (2022-09-27)


### Bug Fixes

* updated tcp listener ([#174](https://github.com/open-feature/flagd/issues/174)) ([b750ed1](https://github.com/open-feature/flagd/commit/b750ed1268b5e6efe779a34e764bad3e781f8e93))

## [0.2.0](https://github.com/open-feature/flagd/compare/v0.1.1...v0.2.0) (2022-09-26)


### ⚠ BREAKING CHANGES

* Updated service to use connect (#163)

### Features

* Updated service to use connect ([#163](https://github.com/open-feature/flagd/issues/163)) ([828d5c4](https://github.com/open-feature/flagd/commit/828d5c4c11157f5b7a77f5041806ba2523186764))


### Bug Fixes

* checkout release tag before running container and binary releases ([#171](https://github.com/open-feature/flagd/issues/171)) ([50fe46f](https://github.com/open-feature/flagd/commit/50fe46fbbf120a0657c1df35b370cdc9051d0f31))

## [0.1.1](https://github.com/open-feature/flagd/compare/v0.1.0...v0.1.1) (2022-09-23)


### Bug Fixes

* bubbles up unclean error exits ([#170](https://github.com/open-feature/flagd/issues/170)) ([9f7db02](https://github.com/open-feature/flagd/commit/9f7db0259d2d24cb880eeddaebd3b8f48758248a))
* upgrade package containing vulnerability ([#162](https://github.com/open-feature/flagd/issues/162)) ([82278c7](https://github.com/open-feature/flagd/commit/82278c7cf08cc6b50f49ab500caf6f9003fc0823))
