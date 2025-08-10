package handler

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestGetNotSafeAddress(t *testing.T) {

	_cookie := "_bl_uid=1wmz8eCq1445tmhU8hktzps2hC51; detect_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyYW5kb21fc3RyIjoiODY3NzM5In0.DaCjSesFMsjGWQkB7iHA1EI5Lp2s3-DTPmxB7nNIPKI; csrftoken=u5xzDP2pcMqbACyYHyVUlJNtmNlr4pIn5i6ullnZNtNunsFbIHvHZk9rteAcyq2l; sessionid=uqs748r6gmq6cjjrqig5461rw8nc3gq9"
	labelAddresses := GetNotSafeAddress("ETH", "0xf510e53ef8da4e45ffa59eb554511a7410e5efd3", _cookie)

	//var labelAddresses LableAddresList
	//err := json.Unmarshal(jsonData, &labelAddresses)
	//if err != nil {
	//	fmt.Println("解析失败:", err)
	//	return
	//}

	if len(labelAddresses.GraphDic.NodeList) > 0 {
		for _, data := range labelAddresses.GraphDic.NodeList {
			fmt.Println("data label: ", data.Label)
			if strings.Contains(data.Label, "huione") {
				fmt.Println("汇旺")
			}
			if strings.Contains(data.Label, "Theft") {
				fmt.Println("盗窃")
			}
			if strings.Contains(data.Label, "Drainer") {
				fmt.Println("诈骗")
			}
			if strings.Contains(data.Label, "Banned") {
				fmt.Println("制裁")
			}
		}
	}

}
func TestGetAddressInfo(t *testing.T) {

	jsonData := []byte(`{
    "success": true,
    "msg": "ok",
    "graph_dic": {
        "node_list": [
            {
                "id": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "USDT Banned Address",
                "title": "0xf510e53ef8da4e45ffa59eb554511a7410e5efd3",
                "layer": 0,
                "addr": "0xf510e53ef8da4e45ffa59eb554511a7410e5efd3",
                "track": "one",
                "pid": 0,
                "color": "#4fae7b",
                "shape": "star",
                "expanded": true,
                "malicious": 1,
                "dex": 0
            },
            {
                "id": "0b632b90-f4b8-4307-9dc9-3fbe36cee590",
                "label": "Inferno Drainer",
                "title": "0xd9779f83632955ba09ca6d53bbb1ff9dbb49448d",
                "layer": 1,
                "addr": "0xd9779f83632955ba09ca6d53bbb1ff9dbb49448d",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "malicious": 1,
                "shape": "circularImage",
                "image": "/malicious.png",
                "dex": 0
            },
            {
                "id": "34d6b952-bcae-467f-9c4c-6042dc26b7ad",
                "label": "0x8cc...04ea7",
                "title": "0x8cc51aabbef95ad9e419d4f10a4988307c404ea7",
                "layer": 1,
                "addr": "0x8cc51aabbef95ad9e419d4f10a4988307c404ea7",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "6702b123-f84b-4aa7-aa9c-0773df8d2ac0",
                "label": "Bridgers.xyz Swap",
                "title": "0xb685760ebd368a891f27ae547391f4e2a289895b",
                "layer": -1,
                "addr": "0xb685760ebd368a891f27ae547391f4e2a289895b",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "exg": 1,
                "shape": "circularImage",
                "image": "/entity.png",
                "dex": 0
            },
            {
                "id": "813789f8-7c30-4bb9-a7a0-27874a088aac",
                "label": "Event: Theft",
                "title": "0xa1d74e57c2f255befe12d9b1d5777871e6aaf547",
                "layer": 1,
                "addr": "0xa1d74e57c2f255befe12d9b1d5777871e6aaf547",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "malicious": 1,
                "shape": "circularImage",
                "image": "/malicious.png",
                "dex": 0
            },
            {
                "id": "b13ad41b-ddf5-4621-a717-8b610ba5c12c",
                "label": "0x03f...21a20",
                "title": "0x03f6c3282c380198f05aa1f51df71eb4d2c21a20",
                "layer": -1,
                "addr": "0x03f6c3282c380198f05aa1f51df71eb4d2c21a20",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "f483289d-758a-4df6-8adf-c320984ce67b",
                "label": "0xf38...14a3b",
                "title": "0xf38c31ee4693597cae748885732f0ffa30314a3b",
                "layer": 1,
                "addr": "0xf38c31ee4693597cae748885732f0ffa30314a3b",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "a4d5e63b-b65c-4238-afa2-8da530c9191a",
                "label": "Event: Theft",
                "title": "0xa1d74e57c2f255befe12d9b1d5777871e6aaf547",
                "layer": -1,
                "addr": "0xa1d74e57c2f255befe12d9b1d5777871e6aaf547",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "malicious": 1,
                "shape": "circularImage",
                "image": "/malicious.png",
                "dex": 0
            },
            {
                "id": "fc918fe0-397d-47a4-9bbc-621d9efdbf7f",
                "label": "0xdb8...b702d",
                "title": "0xdb85e48297d43d6576aaf34f388720dd564b702d",
                "layer": 1,
                "addr": "0xdb85e48297d43d6576aaf34f388720dd564b702d",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "bf906301-572e-4af2-9fc5-6dfa3bf0f58a",
                "label": "0xda3...d72ba",
                "title": "0xda30e0ff85660df8741f9b683add35933c5d72ba",
                "layer": -1,
                "addr": "0xda30e0ff85660df8741f9b683add35933c5d72ba",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "01b27c87-4194-4e5c-a9e9-a594a7ae4e23",
                "label": "0x18a...a850b",
                "title": "0x18a9225871d2654242ed8f6ed7743b56ac7a850b",
                "layer": 1,
                "addr": "0x18a9225871d2654242ed8f6ed7743b56ac7a850b",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "5222acb0-66e6-4da4-aac2-56579429362b",
                "label": "0x5f8...fbf49",
                "title": "0x5f8bce303e9c38b165cee2d238ed254c1ecfbf49",
                "layer": 1,
                "addr": "0x5f8bce303e9c38b165cee2d238ed254c1ecfbf49",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "d7529b52-2310-4ca4-8d33-161de993fe6d",
                "label": "0x682...75b30",
                "title": "0x68268935a31671e12d2c55cb9936fa8114a75b30",
                "layer": -1,
                "addr": "0x68268935a31671e12d2c55cb9936fa8114a75b30",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "edac2457-57d4-40a9-93a3-06489f05ffc5",
                "label": "0x438...7d8aa",
                "title": "0x43820205ccdd444c20d14e1edbe54d9e9347d8aa",
                "layer": 1,
                "addr": "0x43820205ccdd444c20d14e1edbe54d9e9347d8aa",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "6ff97cb9-6aef-43ca-912f-357d934fc81e",
                "label": "Event: Theft",
                "title": "0x12793eec3742b3a64a1465adcb44ecb9318e2868",
                "layer": -1,
                "addr": "0x12793eec3742b3a64a1465adcb44ecb9318e2868",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "malicious": 1,
                "shape": "circularImage",
                "image": "/malicious.png",
                "dex": 0
            },
            {
                "id": "56ebdcf7-fdb4-4c98-a5e2-4f955c4e032f",
                "label": "0x9c8...2d135",
                "title": "0x9c8dae64885cf2497d841dae60b2db00bd02d135",
                "layer": -1,
                "addr": "0x9c8dae64885cf2497d841dae60b2db00bd02d135",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "8dcdce0b-50ce-4a66-b606-52b5d355a586",
                "label": "0xaa5...8f753",
                "title": "0xaa5076095e19027627674c1ea491cbb42a48f753",
                "layer": 1,
                "addr": "0xaa5076095e19027627674c1ea491cbb42a48f753",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "c91763a6-8b0d-4899-a0ee-fd0e7df7f622",
                "label": "0xb95...8f74c",
                "title": "0xb951d1d4103e24d2fd9a7d43f896240f5068f74c",
                "layer": -1,
                "addr": "0xb951d1d4103e24d2fd9a7d43f896240f5068f74c",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "b292d32c-51a3-4916-b73f-3ade2db5a347",
                "label": "0x8c5...f2097",
                "title": "0x8c573edac3e5709781fab23a28a7a0cadd2f2097",
                "layer": 1,
                "addr": "0x8c573edac3e5709781fab23a28a7a0cadd2f2097",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "705be75d-a0cd-4e4b-8c69-75ad84cadf1a",
                "label": "0x8cc...04ea7",
                "title": "0x8cc51aabbef95ad9e419d4f10a4988307c404ea7",
                "layer": -1,
                "addr": "0x8cc51aabbef95ad9e419d4f10a4988307c404ea7",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "37df9210-0f05-4085-a6d4-a1a1954bfe26",
                "label": "0x317...cc6fa",
                "title": "0x31792cb4a4e7425f3385dd3b2f61675edadcc6fa",
                "layer": 1,
                "addr": "0x31792cb4a4e7425f3385dd3b2f61675edadcc6fa",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "8944aef6-62be-4507-abd1-19895e750c4c",
                "label": "0x7cb...ccb6e",
                "title": "0x7cbfabb79dfe80dd78637c6ef74d633e977ccb6e",
                "layer": -1,
                "addr": "0x7cbfabb79dfe80dd78637c6ef74d633e977ccb6e",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "416f1225-3d60-4a7e-afde-69d718e8a744",
                "label": "0x45a...6868d",
                "title": "0x45a7cb67b09e4a4d2647989c797f1c048276868d",
                "layer": -1,
                "addr": "0x45a7cb67b09e4a4d2647989c797f1c048276868d",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "8bbfb0e1-de02-48c9-a809-e2957609e44e",
                "label": "0xda3...d72ba",
                "title": "0xda30e0ff85660df8741f9b683add35933c5d72ba",
                "layer": 1,
                "addr": "0xda30e0ff85660df8741f9b683add35933c5d72ba",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "1baac4b3-0443-4077-98cd-2d6cb6d97600",
                "label": "0x45a...6868d",
                "title": "0x45a7cb67b09e4a4d2647989c797f1c048276868d",
                "layer": 1,
                "addr": "0x45a7cb67b09e4a4d2647989c797f1c048276868d",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "e2b42129-59af-41c3-b514-130f4521375a",
                "label": "0x8c5...f2097",
                "title": "0x8c57891aff90463dbfe9b5c4bf5f75eb493f2097",
                "layer": -1,
                "addr": "0x8c57891aff90463dbfe9b5c4bf5f75eb493f2097",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "8a143748-4d69-442a-9382-600e5a5b330e",
                "label": "0x52d...455de",
                "title": "0x52db548a1519634aec2e48484daf364ac5e455de",
                "layer": -1,
                "addr": "0x52db548a1519634aec2e48484daf364ac5e455de",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "a88ead7b-2988-40bb-8d83-59629ba08484",
                "label": "okx-hot",
                "title": "0xa9ac43f5b5e38155a288d1a01d2cbc4478e14573",
                "layer": -1,
                "addr": "0xa9ac43f5b5e38155a288d1a01d2cbc4478e14573",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "exg": 1,
                "shape": "circularImage",
                "image": "/entity.png",
                "dex": 0
            },
            {
                "id": "9d229993-ae94-4246-89b5-203e1118c39e",
                "label": "0x51b...7f8d1",
                "title": "0x51be3350cdfc71a9b1a5ffbafcca2d9de477f8d1",
                "layer": -1,
                "addr": "0x51be3350cdfc71a9b1a5ffbafcca2d9de477f8d1",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "b3b9f010-32d2-4a1d-94f3-1f054cdd8870",
                "label": "0x28e...48221",
                "title": "0x28ed3639b9358505fe1cf5967d670260cd348221",
                "layer": -1,
                "addr": "0x28ed3639b9358505fe1cf5967d670260cd348221",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "fced5389-9745-484a-bdd6-32c08bb830e8",
                "label": "0x241...e9b3e",
                "title": "0x2411fd5404410a6816627a7c0e60369e7a1e9b3e",
                "layer": -1,
                "addr": "0x2411fd5404410a6816627a7c0e60369e7a1e9b3e",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "d6a6a547-2592-4a9a-be7b-e7dacd3b03a4",
                "label": "0x241...e9b3e",
                "title": "0x2411fd5404410a6816627a7c0e60369e7a1e9b3e",
                "layer": 1,
                "addr": "0x2411fd5404410a6816627a7c0e60369e7a1e9b3e",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "349b465c-7708-41a4-8e05-c7f4019fbc87",
                "label": "0x808...39fa6",
                "title": "0x808abf236ae084a18fbfc17dfa3df7b3a1939fa6",
                "layer": -1,
                "addr": "0x808abf236ae084a18fbfc17dfa3df7b3a1939fa6",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "4c5d4d83-e2e3-48f9-b1d1-825e55aa1d16",
                "label": "0x316...8cc3f",
                "title": "0x316fedd5a12ffeabf5c51b42c411a9cdfc78cc3f",
                "layer": 1,
                "addr": "0x316fedd5a12ffeabf5c51b42c411a9cdfc78cc3f",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "3762654b-27e7-4dd9-adb6-356ff468051b",
                "label": "0x316...8cc3f",
                "title": "0x316fedd5a12ffeabf5c51b42c411a9cdfc78cc3f",
                "layer": -1,
                "addr": "0x316fedd5a12ffeabf5c51b42c411a9cdfc78cc3f",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "6f8b2865-87f0-4a5f-bfa1-5c2818208a40",
                "label": "0xbb5...77ae3",
                "title": "0xbb5b2d5b59c889b83493de5fe16cd3e82fd77ae3",
                "layer": -1,
                "addr": "0xbb5b2d5b59c889b83493de5fe16cd3e82fd77ae3",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "78436200-4927-4006-91a6-653b2c178621",
                "label": "0x4e6...5becb",
                "title": "0x4e66c907a747aec33acc65d9a9f6d5eb2625becb",
                "layer": -1,
                "addr": "0x4e66c907a747aec33acc65d9a9f6d5eb2625becb",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "8435fb7b-7878-42e6-8f37-ad87ea66d6be",
                "label": "okx-hot",
                "title": "0xf7858da8a6617f7c6d0ff2bcafdb6d2eedf64840",
                "layer": -1,
                "addr": "0xf7858da8a6617f7c6d0ff2bcafdb6d2eedf64840",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "exg": 1,
                "shape": "circularImage",
                "image": "/entity.png",
                "dex": 0
            },
            {
                "id": "2af28f8d-a08f-4fce-8f59-f4557cf66e8b",
                "label": "0x91e...895a5",
                "title": "0x91efd4877e7c3a7815d3f9533a20f64437b895a5",
                "layer": 1,
                "addr": "0x91efd4877e7c3a7815d3f9533a20f64437b895a5",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "eb6ff87d-c116-40e3-9c30-109cf01698b3",
                "label": "0x100...8a040",
                "title": "0x10066bb0bc1127964b9877419e3a00572a98a040",
                "layer": -1,
                "addr": "0x10066bb0bc1127964b9877419e3a00572a98a040",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "a25b527c-e911-4ae4-a531-40f3eb6277f3",
                "label": "0x0ed...dc60e",
                "title": "0x0ede4d33afa4ca1a1d85e61a0e6dd783d9edc60e",
                "layer": 1,
                "addr": "0x0ede4d33afa4ca1a1d85e61a0e6dd783d9edc60e",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "47843943-cb8e-48ed-a850-655ce7711ca9",
                "label": "[Contract] 0x7ae...d3677",
                "title": "0x7ae91b984da4795d9fd88419d051d0842f8d3677",
                "layer": -1,
                "addr": "0x7ae91b984da4795d9fd88419d051d0842f8d3677",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "contract": 1,
                "dex": 0
            },
            {
                "id": "2a180f1c-78d9-4dc0-b30f-36b354c53787",
                "label": "0x91e...895a5",
                "title": "0x91efd4877e7c3a7815d3f9533a20f64437b895a5",
                "layer": -1,
                "addr": "0x91efd4877e7c3a7815d3f9533a20f64437b895a5",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "466833c3-5332-40ff-9230-4c484f9bc1ca",
                "label": "0x095...960d4",
                "title": "0x095cbe41e11838caf04e36f1f57339a1792960d4",
                "layer": -1,
                "addr": "0x095cbe41e11838caf04e36f1f57339a1792960d4",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "a31807b0-710e-4256-a40a-25f35f602019",
                "label": "kraken",
                "title": "0x6cd4fb496909a3a01c8dcffaccfd54984f291a25",
                "layer": 1,
                "addr": "0x6cd4fb496909a3a01c8dcffaccfd54984f291a25",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "exg": 1,
                "shape": "circularImage",
                "image": "/entity.png",
                "dex": 0
            },
            {
                "id": "9853d89d-85f4-41df-9e2f-bf0dbcad3d9c",
                "label": "0x167...98661",
                "title": "0x167908e85a7a88814d35f5dd827c782a02698661",
                "layer": 1,
                "addr": "0x167908e85a7a88814d35f5dd827c782a02698661",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "b1e14a67-fdc1-4103-9ca5-ad9ba979e561",
                "label": "0x823...da10f",
                "title": "0x823e3ae9b9f0edc07710dd029241cc43668da10f",
                "layer": -1,
                "addr": "0x823e3ae9b9f0edc07710dd029241cc43668da10f",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "543a26b2-e206-4794-9cc0-45761c165a9a",
                "label": "0x80a...02f7b",
                "title": "0x80a7323af41d479356a5478680008bef8a702f7b",
                "layer": 1,
                "addr": "0x80a7323af41d479356a5478680008bef8a702f7b",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "aab7135e-b03b-401d-bafc-fb716aff0bff",
                "label": "0xece...bf54e",
                "title": "0xece6d28f9e13e9602110c1fca9e2de13c07bf54e",
                "layer": -1,
                "addr": "0xece6d28f9e13e9602110c1fca9e2de13c07bf54e",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "fcbd92f2-3cf4-485e-aa9b-cddf74779ac8",
                "label": "0x674...b12ee",
                "title": "0x6744807e3f6e9d1a547ccd1493feb58c020b12ee",
                "layer": -1,
                "addr": "0x6744807e3f6e9d1a547ccd1493feb58c020b12ee",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "94615aed-467a-4105-9925-882e6d281fa9",
                "label": "0x32c...0ea90",
                "title": "0x32cc1325b5976620e87464f61fad9e968ed0ea90",
                "layer": -1,
                "addr": "0x32cc1325b5976620e87464f61fad9e968ed0ea90",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "b3d4caf6-4904-4930-8ac9-a042a9449bc4",
                "label": "0xf08...850c1",
                "title": "0xf08fa5e7116f7679261e571d1d24edd3f0d850c1",
                "layer": 1,
                "addr": "0xf08fa5e7116f7679261e571d1d24edd3f0d850c1",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "1842f98c-47f4-4409-a756-66277dd3a0c1",
                "label": "0x0ed...dc60e",
                "title": "0x0ede4d33afa4ca1a1d85e61a0e6dd783d9edc60e",
                "layer": -1,
                "addr": "0x0ede4d33afa4ca1a1d85e61a0e6dd783d9edc60e",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "b9e1b93d-808b-40bc-911e-e1e35aa64ab2",
                "label": "0x5b9...496cc",
                "title": "0x5b90abc7361c3b4a119681462398ba41c37496cc",
                "layer": -1,
                "addr": "0x5b90abc7361c3b4a119681462398ba41c37496cc",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "c6d3e28d-768d-4442-b47c-492e26154c83",
                "label": "0xdd7...ce6ef",
                "title": "0xdd7905b17ffe3e1b37b4550e5649cdb9f4ece6ef",
                "layer": -1,
                "addr": "0xdd7905b17ffe3e1b37b4550e5649cdb9f4ece6ef",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "7af9790b-eda7-4efe-8ab7-a542b0ff4d40",
                "label": "0xdd3...9f015",
                "title": "0xdd33b3860c5fcec7df8221c2e8030cf750b9f015",
                "layer": -1,
                "addr": "0xdd33b3860c5fcec7df8221c2e8030cf750b9f015",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "84109363-dc7b-4fd7-aaa0-57c231541121",
                "label": "0x386...9e7dd",
                "title": "0x386c398e721bc92931f4f596ff424d672609e7dd",
                "layer": -1,
                "addr": "0x386c398e721bc92931f4f596ff424d672609e7dd",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "bed23184-5c45-470e-83dc-6144d2393b9e",
                "label": "0x4c8...99999",
                "title": "0x4c8248104d05900876e95adf4bd7a20129999999",
                "layer": -1,
                "addr": "0x4c8248104d05900876e95adf4bd7a20129999999",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "d4ba3b04-d36d-4178-aa6b-c6ed0a1be828",
                "label": "0x4c8...99999",
                "title": "0x4c8248104d05900876e95adf4bd7a20129999999",
                "layer": 1,
                "addr": "0x4c8248104d05900876e95adf4bd7a20129999999",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "0912a14f-f011-4c82-b769-e044ae0441cf",
                "label": "0x012...eed0c",
                "title": "0x012746bef071aedcd10a0f6f54ccac13b56eed0c",
                "layer": 1,
                "addr": "0x012746bef071aedcd10a0f6f54ccac13b56eed0c",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "fa000a3e-e24b-4cbf-afba-4f00de45963f",
                "label": "0x0a9...36dd0",
                "title": "0x0a98c712219662abf5114ed324701df865e36dd0",
                "layer": -1,
                "addr": "0x0a98c712219662abf5114ed324701df865e36dd0",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "14f07197-0ebb-4b12-8499-000492964874",
                "label": "0xc6b...6b60e",
                "title": "0xc6b7aa092978fd4bd39e24ff58b9e181f786b60e",
                "layer": 1,
                "addr": "0xc6b7aa092978fd4bd39e24ff58b9e181f786b60e",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "a61d54fb-8890-4f8d-a487-bcb401de43cd",
                "label": "0x0df...b5feb",
                "title": "0x0df127a9c25f6e209901a2564363b59d28ab5feb",
                "layer": -1,
                "addr": "0x0df127a9c25f6e209901a2564363b59d28ab5feb",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "dbecbba4-00ff-4053-b8b4-f20737a5c65d",
                "label": "0x781...8a794",
                "title": "0x781292ab0a241b98b05f0d8c23ba2c2bc588a794",
                "layer": -1,
                "addr": "0x781292ab0a241b98b05f0d8c23ba2c2bc588a794",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "a3b0feef-4c98-42e9-bb0f-19f1054b9648",
                "label": "0x6e9...01f88",
                "title": "0x6e9cce3bd8137aece90248b529be91db45901f88",
                "layer": -1,
                "addr": "0x6e9cce3bd8137aece90248b529be91db45901f88",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "a720f9a9-ae5e-4cb6-96ce-b7d621b7e4bc",
                "label": "0xd18...f03c6",
                "title": "0xd18c2028446ef51daf48455f10572cb9888f03c6",
                "layer": -1,
                "addr": "0xd18c2028446ef51daf48455f10572cb9888f03c6",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "b4f90922-170c-4034-b33a-11da69144049",
                "label": "0x6b1...1e82e",
                "title": "0x6b1a1ae3357843ae003c52cffc5e0a3d35a1e82e",
                "layer": -1,
                "addr": "0x6b1a1ae3357843ae003c52cffc5e0a3d35a1e82e",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "c793fd3a-63b6-4996-82d9-070882052f56",
                "label": "0xe5b...f0b46",
                "title": "0xe5bf5f7c576433c63d6f93baffd0f2b2ee5f0b46",
                "layer": -1,
                "addr": "0xe5bf5f7c576433c63d6f93baffd0f2b2ee5f0b46",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "eb86cc20-5e7d-4bba-9f98-0cf4546590e1",
                "label": "0x8d9...89396",
                "title": "0x8d9fa0bca2efe7461467f1a6a6464eab1ab89396",
                "layer": 1,
                "addr": "0x8d9fa0bca2efe7461467f1a6a6464eab1ab89396",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "cef9001b-8449-4db3-a6c9-19d108d5bff7",
                "label": "united-states-dollar.eth",
                "title": "0xffa52ef9639865deaaa21f76761373bc3b9b601f",
                "layer": -1,
                "addr": "0xffa52ef9639865deaaa21f76761373bc3b9b601f",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "ens": 1,
                "dex": 0
            },
            {
                "id": "c2d7232b-d64d-4706-9174-77e5287948fd",
                "label": "Bridgers.xyz Swap",
                "title": "0x92e929d8b2c8430bcaf4cd87654789578bb2b786",
                "layer": -1,
                "addr": "0x92e929d8b2c8430bcaf4cd87654789578bb2b786",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "exg": 1,
                "shape": "circularImage",
                "image": "/entity.png",
                "dex": 0
            },
            {
                "id": "2e6e7fd6-f066-4ea9-ae23-539a84d3fe8b",
                "label": "[Contract] 0x103...4ffc5",
                "title": "0x10374bc1c4ca086e22673dbc4d702fee74c4ffc5",
                "layer": -1,
                "addr": "0x10374bc1c4ca086e22673dbc4d702fee74c4ffc5",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "contract": 1,
                "dex": 0
            },
            {
                "id": "40bd1521-83bb-445b-b124-a0e5dada0d4f",
                "label": "Event: Phishing",
                "title": "0x77e7c5cbeaad915cf5462064b02984e16a902e67",
                "layer": -1,
                "addr": "0x77e7c5cbeaad915cf5462064b02984e16a902e67",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "malicious": 1,
                "shape": "circularImage",
                "image": "/malicious.png",
                "dex": 0
            },
            {
                "id": "5809ed55-956d-48b7-92c1-834799eaaf3d",
                "label": "0xbcf...84af1",
                "title": "0xbcfe78cd98f0c55eb5decd714800038c11884af1",
                "layer": -1,
                "addr": "0xbcfe78cd98f0c55eb5decd714800038c11884af1",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "d9d6dcc5-ac3f-49ff-b0f6-b7ad57afb5d2",
                "label": "Event: Theft",
                "title": "0xfb384e2153a3322d0a69d716503d6f65d47bbc93",
                "layer": -1,
                "addr": "0xfb384e2153a3322d0a69d716503d6f65d47bbc93",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "malicious": 1,
                "shape": "circularImage",
                "image": "/malicious.png",
                "dex": 0
            },
            {
                "id": "6282668b-b44a-452b-8e01-59c0f9ed618b",
                "label": "0xe5b...f0b46",
                "title": "0xe5bf5f7c576433c63d6f93baffd0f2b2ee5f0b46",
                "layer": 1,
                "addr": "0xe5bf5f7c576433c63d6f93baffd0f2b2ee5f0b46",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "c9b2a0a0-4b90-44af-bf55-7d27614d4d09",
                "label": "0x577...09d3d",
                "title": "0x5776552dcfe8605a8efa8036b9ad74ee10109d3d",
                "layer": -1,
                "addr": "0x5776552dcfe8605a8efa8036b9ad74ee10109d3d",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "a286cea3-8a1d-4a2c-bf12-a0cfeb0034ac",
                "label": "0x727...b1c33",
                "title": "0x727e3a72b0014dab1c2f8cb52f63dee798fb1c33",
                "layer": -1,
                "addr": "0x727e3a72b0014dab1c2f8cb52f63dee798fb1c33",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "c6a638ba-5c20-44c0-9a99-22add03bb1ff",
                "label": "0x2d8...aa5b8",
                "title": "0x2d8cfacd4651c255efc05150d798dd88975aa5b8",
                "layer": 1,
                "addr": "0x2d8cfacd4651c255efc05150d798dd88975aa5b8",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "8b318e8d-4481-4259-8b6c-841b6f93619d",
                "label": "0x439...192aa",
                "title": "0x4393ca1938b6fa946d0c01ba8af256c03cc192aa",
                "layer": -1,
                "addr": "0x4393ca1938b6fa946d0c01ba8af256c03cc192aa",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "ffcba112-b1be-426c-972d-c60b8de6801c",
                "label": "0x2d8...aa5b8",
                "title": "0x2d8cfacd4651c255efc05150d798dd88975aa5b8",
                "layer": -1,
                "addr": "0x2d8cfacd4651c255efc05150d798dd88975aa5b8",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "b61d25a1-ec0a-4fd2-83ec-80ccd9cb9047",
                "label": "0x125...1cc0a",
                "title": "0x125aa21ebaecbf5a76718ef7d15f31b858f1cc0a",
                "layer": -1,
                "addr": "0x125aa21ebaecbf5a76718ef7d15f31b858f1cc0a",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "ad8628e9-8640-4759-b974-83f5f6290aa4",
                "label": "Inferno Drainer",
                "title": "0xd9779f83632955ba09ca6d53bbb1ff9dbb49448d",
                "layer": -1,
                "addr": "0xd9779f83632955ba09ca6d53bbb1ff9dbb49448d",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "malicious": 1,
                "shape": "circularImage",
                "image": "/malicious.png",
                "dex": 0
            },
            {
                "id": "4c8ed30e-216d-4831-be07-835049d34c59",
                "label": "0x348...fd354",
                "title": "0x348358bbdfd7f5b7463f98bdd6e309955c1fd354",
                "layer": 1,
                "addr": "0x348358bbdfd7f5b7463f98bdd6e309955c1fd354",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "e401dafa-19b6-4cae-a1cc-b98acaf18d6d",
                "label": "0xa2f...f2208",
                "title": "0xa2f929870e7ff83d209617134bd5816bd54f2208",
                "layer": -1,
                "addr": "0xa2f929870e7ff83d209617134bd5816bd54f2208",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "0f3e5735-f82f-4858-827b-6b0cb8fa84e7",
                "label": "0x02c...94aa7",
                "title": "0x02cc9077d637c7ddb36b8a70f4d04f4ad1694aa7",
                "layer": 1,
                "addr": "0x02cc9077d637c7ddb36b8a70f4d04f4ad1694aa7",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "06326d94-a1a6-4233-a0dc-0bdc68319ec6",
                "label": "0xd92...366a9",
                "title": "0xd925c1897a7046822c11c235d8a730459c6366a9",
                "layer": -1,
                "addr": "0xd925c1897a7046822c11c235d8a730459c6366a9",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "31e47172-12ce-4729-90f7-5dfdb44bada6",
                "label": "0x577...09d3d",
                "title": "0x5776552dcfe8605a8efa8036b9ad74ee10109d3d",
                "layer": 1,
                "addr": "0x5776552dcfe8605a8efa8036b9ad74ee10109d3d",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "2ad7218f-296f-4c10-aa61-f87c4a30ac25",
                "label": "0x06a...4ee23",
                "title": "0x06a4c12f65762320cff3ed70b5e2089180c4ee23",
                "layer": -1,
                "addr": "0x06a4c12f65762320cff3ed70b5e2089180c4ee23",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "e8bbf4e7-e0a5-4f16-bc76-ef461c082a59",
                "label": "0xf58...d0061",
                "title": "0xf584290b52f2f330edb8323b7851df8c5d4d0061",
                "layer": -1,
                "addr": "0xf584290b52f2f330edb8323b7851df8c5d4d0061",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "f137ffd5-b153-420e-87a8-edc69b39d751",
                "label": "0x98b...2f402",
                "title": "0x98bc5385a568f8a6466d71456712272b8532f402",
                "layer": 1,
                "addr": "0x98bc5385a568f8a6466d71456712272b8532f402",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "61cc42a0-285e-421b-82e7-c699cf9adf2a",
                "label": "0x98b...2f402",
                "title": "0x98bc5385a568f8a6466d71456712272b8532f402",
                "layer": -1,
                "addr": "0x98bc5385a568f8a6466d71456712272b8532f402",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "bcfcbda8-114d-4911-97b0-bae644b10805",
                "label": "0xc04...59b6d",
                "title": "0xc04e57c1ee4044b6b29f79ac26c15e600f359b6d",
                "layer": -1,
                "addr": "0xc04e57c1ee4044b6b29f79ac26c15e600f359b6d",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "56b41b43-8058-49c2-80a7-ad3d25e51ad4",
                "label": "0xf98...f58fc",
                "title": "0xf9841b75df4643a7357ddcff552cc474edbf58fc",
                "layer": -1,
                "addr": "0xf9841b75df4643a7357ddcff552cc474edbf58fc",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "100150cf-61ff-4b0c-b328-480730e8301f",
                "label": "0xe0d...3829d",
                "title": "0xe0d8e2541a23b07091894cd7fa7ce57e33a3829d",
                "layer": -1,
                "addr": "0xe0d8e2541a23b07091894cd7fa7ce57e33a3829d",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "03229d72-5ad1-4815-9bbc-19d23fe49cbe",
                "label": "0xcf4...d195a",
                "title": "0xcf412e0e95a56df7afcecd17d252c650239d195a",
                "layer": -1,
                "addr": "0xcf412e0e95a56df7afcecd17d252c650239d195a",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "826e97d9-fdbd-4100-b8e3-a34e65937f06",
                "label": "0xc04...59b6d",
                "title": "0xc04e57c1ee4044b6b29f79ac26c15e600f359b6d",
                "layer": 1,
                "addr": "0xc04e57c1ee4044b6b29f79ac26c15e600f359b6d",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "2288997a-6ca3-4a13-a1f4-282a1ff7bf49",
                "label": "0xbfe...df2a5",
                "title": "0xbfe2fb47a548a49b8615407e0ae858fb8c3df2a5",
                "layer": 1,
                "addr": "0xbfe2fb47a548a49b8615407e0ae858fb8c3df2a5",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "6e4ee2d5-9aca-4f20-8e81-0cda0c57900e",
                "label": "[Contract] 0xf43...56561",
                "title": "0xf43dff8f7dc04dd98f5ae6c9a9273e0bd3c56561",
                "layer": 1,
                "addr": "0xf43dff8f7dc04dd98f5ae6c9a9273e0bd3c56561",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "contract": 1,
                "dex": 0
            },
            {
                "id": "dc15b9cb-0f42-42e8-ad3c-6cd4c523eae7",
                "label": "0x978...a794b",
                "title": "0x978b4f5e215f3779a31f0bb3a7e65181b8ba794b",
                "layer": 1,
                "addr": "0x978b4f5e215f3779a31f0bb3a7e65181b8ba794b",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "21898669-9a6d-4d7a-b49c-1a1b45798ae5",
                "label": "0xdbf...1b497",
                "title": "0xdbf5102a452f753df2ff2e9b835fc5d2d121b497",
                "layer": 1,
                "addr": "0xdbf5102a452f753df2ff2e9b835fc5d2d121b497",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "2240c595-f899-41c0-afc0-c7e0556c5fe8",
                "label": "0x0df...b5feb",
                "title": "0x0df127a9c25f6e209901a2564363b59d28ab5feb",
                "layer": 1,
                "addr": "0x0df127a9c25f6e209901a2564363b59d28ab5feb",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "e729e355-cc6b-429c-bcd6-d4373baa3b0f",
                "label": "0x140...d1b12",
                "title": "0x140a15c8d14335d68138dc2ea5358e48116d1b12",
                "layer": 1,
                "addr": "0x140a15c8d14335d68138dc2ea5358e48116d1b12",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "d1936264-6221-45f4-b1f4-2bbf5adf74e4",
                "label": "0xaab...44938",
                "title": "0xaabb7e129cb3f3ec4e834a7c515d02cfc1d44938",
                "layer": -1,
                "addr": "0xaabb7e129cb3f3ec4e834a7c515d02cfc1d44938",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            },
            {
                "id": "c18e8de7-b5c9-49b9-a92f-e7d3693ac6da",
                "label": "0x524...c0f3c",
                "title": "0x524e7b13c1bc2b3c4706e77e6ac30f858e3c0f3c",
                "layer": -1,
                "addr": "0x524e7b13c1bc2b3c4706e77e6ac30f858e3c0f3c",
                "track": "one",
                "pid": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "expanded": false,
                "dex": 0
            }
        ],
        "edge_list": [
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "0b632b90-f4b8-4307-9dc9-3fbe36cee590",
                "label": "0.0442 ETH",
                "val": 0.0442,
                "tx_time": "Aug 09, 11:01 AM",
                "tx_hash_list": [
                    "0x14ec1892b9593fc3d35c049576d6374d2f0e6ec5902183819fa027ec41873239",
                    "0x91028dce2f575aa31bcdf0826d213effcf669951456a266d48f7990f3195a113",
                    "0xe8bac8a82c8fb22e98e00b8ae3370e192f67057eac52a8b5c1a9e5890b5af48f",
                    "0x9b61addb3ba42b990d30f8f62558cea6540841e9c2399caab0cabd9fbedf12fe",
                    "0x44b6d82784d04ae269662eae26485a6505ccfceaeeeed47993cd174526ac4769",
                    "0x7a68ab75ead825387deef0e322e443edfc60947efe3685c15d0477ee5536d825"
                ]
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "34d6b952-bcae-467f-9c4c-6042dc26b7ad",
                "label": "0.0063 ETH",
                "val": 0.0063,
                "tx_time": "Aug 09, 10:59 AM",
                "tx_hash_list": [
                    "0xed714c1238c1fec0cede8cf6b69a3d56ce1329f9a16961a68ecba3920a60670b",
                    "0x94d6eeabdd0ce5b4ac5b0d140db753cd1cbdaaaaf9e834db7170d50972ea008c"
                ]
            },
            {
                "from": "6702b123-f84b-4aa7-aa9c-0773df8d2ac0",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0069 ETH",
                "val": 0.0069,
                "tx_hash_list": [
                    "0xd2ae38c8095fd09a7d8de231b53e21db208be97ab2deda82bb41bc3b22a8ceb2"
                ],
                "tx_time": "Aug 09, 10:58 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "813789f8-7c30-4bb9-a7a0-27874a088aac",
                "label": "0.01 ETH",
                "val": 0.01,
                "tx_time": "Aug 09, 10:57 AM",
                "tx_hash_list": [
                    "0x31ce76867ea38037939430a03bb8cd920c7f6135c9cb0ed1ead487ae97a42026"
                ]
            },
            {
                "from": "b13ad41b-ddf5-4621-a717-8b610ba5c12c",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.01 ETH",
                "val": 0.01,
                "tx_hash_list": [
                    "0x9f1a5ab8e305e6fe95f1960d1af4ea49c3f20e1e78df3fa28450a6338c71c3aa"
                ],
                "tx_time": "Aug 09, 10:56 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "f483289d-758a-4df6-8adf-c320984ce67b",
                "label": "0.0009 ETH",
                "val": 0.0009,
                "tx_time": "Aug 07, 09:26 AM",
                "tx_hash_list": [
                    "0x39b9ad1807c2a192a09ada27a2d78ce0e123d82c83b01d01114e59296e8ddc75"
                ]
            },
            {
                "from": "a4d5e63b-b65c-4238-afa2-8da530c9191a",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.001 ETH",
                "val": 0.001,
                "tx_hash_list": [
                    "0x11aabf5c112055afbed5bfff7d7daa928d071d985aae149253d2371022aeee17"
                ],
                "tx_time": "Aug 07, 09:22 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "fc918fe0-397d-47a4-9bbc-621d9efdbf7f",
                "label": "0.0015 ETH",
                "val": 0.0015,
                "tx_time": "Jul 28, 05:59 PM",
                "tx_hash_list": [
                    "0x2bc45300085c9430aeee73559d7c27b80977b96d0b47a4eb786d21e2cdccc59e",
                    "0x25095b561e186975f6a2dc8e6a497cf577bf3816d4d52d1b0a2929ff9957c5fd",
                    "0x5ba5b0d65f162c1452f5c339a1444d9f5461a3c5f460c662feffaf13d5c4f27d",
                    "0x1bfc52078a81d8f910ffc7b4f9a829cd43d1c4384e6f6f9b67d6d27eb5a27cf5",
                    "0xaabcf9a8665e6379180fc80e5c1891c3a656c0a93d2074c79239885585426cb3",
                    "0x8040eae4d9b2d38219c840b280ea638f1cb1e1786487ed415e3a5fda7e143614",
                    "0x5bd732aa81af5de2213454feb4c1232a8e8bf54f05de86a749e15e7c2c6d8016"
                ]
            },
            {
                "from": "bf906301-572e-4af2-9fc5-6dfa3bf0f58a",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0082 ETH",
                "val": 0.0082,
                "tx_hash_list": [
                    "0x645115357ead77da4b6b26c84a000f7ddcd1803aedc1946842f18ad0218532e0",
                    "0xefaf671ad2a9b3b0ff0ebec6418132d6d079e997f721cfa515a8aa909fa1a492",
                    "0xdc70f0cae1126e6357b203f21df3a96a7bda2ff5bde284e2ef0b4b90a57c3364",
                    "0xa85e9bfd986e1e6cd328ffa138adf3f2dd934b16c151b7c034040de3f55c06c4",
                    "0xaf61200e5d11f20bdd6736190690dce3e7f3dd4b248212d1886b1a7bd07ea958",
                    "0xa10da884d9c67594449e43d74f66c264099d084eeedb83fd2857662ebd9b0461",
                    "0xa96d999cd96e39ca874082b78b0dd22c27fecb85bfdcc4794d15b13e540e0000",
                    "0x5fbd3c5de9280a53f22b6d742355c91c94c8a00f977cf23a292fb495a2c0820c",
                    "0xcbf65841c1fd4c29931efe66611f83893f4e0280cfee2136d361e35ce650a805",
                    "0x2e2e61bbef51902b25c850aab655384962baf24b29dfb6bf1bbe108421040bfd",
                    "0xd444f9e1a71c25da1cef5a938634ee0784839d9a51f03160ca68503c2f7eccf2",
                    "0x26f9ab03b7c967c3afbfa9d4e97573abd61fdb98d2202c25c29747e693de8e6a",
                    "0x23a309033461de00d8af46547a6fcecfd11bce0c15c401d92fd9c13c6c57081f"
                ],
                "tx_time": "Jul 24, 02:54 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "01b27c87-4194-4e5c-a9e9-a594a7ae4e23",
                "label": "0.0002 ETH",
                "val": 0.0002,
                "tx_time": "Jul 22, 07:08 PM",
                "tx_hash_list": [
                    "0x037d57a3e8f8c926221865ec33cb56f8277cb060d1b62063bd5c0099fb21d31e"
                ]
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "5222acb0-66e6-4da4-aac2-56579429362b",
                "label": "0.01 ETH",
                "val": 0.01,
                "tx_time": "Jul 22, 07:06 PM",
                "tx_hash_list": [
                    "0x449655ee7d6df02d6e0df5fa4174d0e6d5d7b7c3750590a3ade18e5ae6057257"
                ]
            },
            {
                "from": "d7529b52-2310-4ca4-8d33-161de993fe6d",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.011 ETH",
                "val": 0.011,
                "tx_hash_list": [
                    "0x7b4189db4238cc6056ea5a3b007c55e4bbfc233e63d884c2debeb38ce39978e4"
                ],
                "tx_time": "Jul 22, 07:04 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "edac2457-57d4-40a9-93a3-06489f05ffc5",
                "label": "0.0034 ETH",
                "val": 0.0034,
                "tx_time": "Jul 14, 03:27 PM",
                "tx_hash_list": [
                    "0xb836617bd88753da38288e1b8de91ab07e3b906acf5db2fdb8c90a82b25295f0",
                    "0xd14832be6abc35adce40798cf9a83a64a5541e04570f52d28628cfcba9432020"
                ]
            },
            {
                "from": "6ff97cb9-6aef-43ca-912f-357d934fc81e",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.003 ETH",
                "val": 0.003,
                "tx_hash_list": [
                    "0x508f5894dd0ceadf9dfaeb9a96a044128ffae2863aeb72afdf16f6191cbce4e9"
                ],
                "tx_time": "Jul 14, 03:26 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "56ebdcf7-fdb4-4c98-a5e2-4f955c4e032f",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.001 ETH",
                "val": 0.001,
                "tx_hash_list": [
                    "0x7008e78c63bcea2e8d7668f17dcd1f64eae3930d2687ed541d5294e6919fc2dc"
                ],
                "tx_time": "Jul 02, 03:02 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "8dcdce0b-50ce-4a66-b606-52b5d355a586",
                "label": "0.0451 ETH",
                "val": 0.0451,
                "tx_time": "Jul 02, 12:05 AM",
                "tx_hash_list": [
                    "0x55ea4945bcc41305eb510fd6157866645e7abeb9cb69896691eae0a2e703b3f8",
                    "0x200c90ffa0f40b7c446d6f3db995ae551d98f92fcd0f376df0deaa95926a623f",
                    "0xeaf70a68e42adb9d3230a58002426bee27a96c08f73f6e23b9b6ae7d76d8447f",
                    "0xf9ed36f93a80a81fcc68befb57031a1255437ef9afcd7dd3f6dc4be9ead2e677"
                ]
            },
            {
                "from": "c91763a6-8b0d-4899-a0ee-fd0e7df7f622",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.01 ETH",
                "val": 0.01,
                "tx_hash_list": [
                    "0x39bb8882158b59d560168b7464c2f0793efdc0f6aacdd80f62b01e68624f5ff9"
                ],
                "tx_time": "Jul 02, 12:03 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "b292d32c-51a3-4916-b73f-3ade2db5a347",
                "label": "0.0664 ETH",
                "val": 0.0664,
                "tx_time": "Jun 23, 08:15 PM",
                "tx_hash_list": [
                    "0xa7b856a7a819129540e754f8789d7a4262136703439a380264a8d51bb7cb6ad2",
                    "0xbf30f8899cc220dad69471a3c7a8ef81f45d04a61e74c20e6952a07776e02989",
                    "0xac6441b696c4fa6cdf8113849a378a3b1804775506545b87352f445859c20c67",
                    "0xab6c1421bd15e63fc714633576de328ca262e4ab78f785d1ba8470ea234f9e92",
                    "0xfc815d76fd2218370e28fb6790869fc6532a0229b4c1d9b56c0dabd453946473",
                    "0x8dc0aa2eb83eb1c333122a7be77790154ab514c48c32e6f504f8764fa4445d97",
                    "0x6372bac9797c9ed9fa3fa9d81a63c87fc95cd832a64c5d6363e4fc408993fe11",
                    "0x9cad329de34f09f81830ba5c548f4297b49eebba3f0544772d436d55f2bdd5d0",
                    "0xf3d418a967de081160642cee39bd549222c1a79a0db93004d697c300c1a81500",
                    "0x41f1e5aec71a093762ef01e60526d2e723f0c575f818bf45aaff214a3608b8d3",
                    "0x237c6a95d69e9016b3d20b7f55410c1c3011ff3ea83b6861f1774a3471a90b46",
                    "0xdacd2c800231d5492313864a14762410ff37c35a5f6de196aebbf438953551d6"
                ]
            },
            {
                "from": "705be75d-a0cd-4e4b-8c69-75ad84cadf1a",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0007 ETH",
                "val": 0.0007,
                "tx_hash_list": [
                    "0xa1a747f395ada76618441236e88242299227bd611535588dc501024c144ae353"
                ],
                "tx_time": "Jun 23, 08:01 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "37df9210-0f05-4085-a6d4-a1a1954bfe26",
                "label": "0.0217 ETH",
                "val": 0.0217,
                "tx_time": "Jun 23, 04:55 PM",
                "tx_hash_list": [
                    "0x8dfd55b5c2f666c645a31bf076932ad90d15b84eb814369db37d8fb3d5a63f84",
                    "0xee02ed00c8178b77066de941b2c58bcc754a7e6b762ec595d0f131c28f54bec5",
                    "0x9f931641ac0b8d54bb146f7bff197d4aef9568617e37cf489a2bdce6da828bfd"
                ]
            },
            {
                "from": "8944aef6-62be-4507-abd1-19895e750c4c",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0089 ETH",
                "val": 0.0089,
                "tx_hash_list": [
                    "0x6e79b53930658e77528c93efe7c6ffb88c08849860088b2da63b05538fbe2ac2",
                    "0x934dee6de81ae10e503513e3fcb930b43884ab5b8f68f67416c46d3df939a643"
                ],
                "tx_time": "Jun 23, 03:05 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "416f1225-3d60-4a7e-afde-69d718e8a744",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0044 ETH",
                "val": 0.0044,
                "tx_hash_list": [
                    "0x1b3747f6ffb13d48e4c4ebaa8582bcabd6471bd26bff29aadd76ea7032acef64",
                    "0x44937b7f07d594a8d20ebf759ed0a73ecdba44fa188eed6bfcbf6d1a4970d3dd",
                    "0x6625f93ecd0503c10a56e7990b44c2ca56812da555ad206151977655449dfe9b",
                    "0xfa9fc14ad514e2fb5883581e49f3870fa335fba5f0f2aedff24c636845184854",
                    "0x21b4edc1a81394f9f91b8965206e62bda3912dd565a6aab7ac8ef4dadbb68789",
                    "0xc17f05c44ed0216c94b75b2912c1ae2879a11c6e474822a4187438bb279410b3",
                    "0x2c4a863f2a11a621b172667404f36588d27f7b665a4b3ae40c6b6d5e47733259"
                ],
                "tx_time": "Jun 23, 09:39 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "8bbfb0e1-de02-48c9-a809-e2957609e44e",
                "label": "0.0003 ETH",
                "val": 0.0003,
                "tx_time": "Jun 23, 09:27 AM",
                "tx_hash_list": [
                    "0xe6b54df6ef2d977f7d85016326ee55a01b475154ea02da1d94a939fea169ccc1"
                ]
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "1baac4b3-0443-4077-98cd-2d6cb6d97600",
                "label": "0.001 ETH",
                "val": 0.001,
                "tx_time": "Jun 23, 07:42 AM",
                "tx_hash_list": [
                    "0x0a82a002f4f9e17b8141d2148c2039edad2672cbd003f8bbe03ac70f8cb3f18e"
                ]
            },
            {
                "from": "e2b42129-59af-41c3-b514-130f4521375a",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0 ETH",
                "val": 0.0,
                "tx_hash_list": [
                    "0xdee3a5264ed571e02590f2e93a160e34eeee04907bef9998a8cff7f4a52f2bb0",
                    "0xc95fed8a27936d95a022054e50662712fb337b366261501665016c5cf2cdcfc7"
                ],
                "tx_time": "Jun 23, 07:07 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "8a143748-4d69-442a-9382-600e5a5b330e",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0166 ETH",
                "val": 0.0166,
                "tx_hash_list": [
                    "0x83a41d1636ad25a69734f70c81b898eee962e975dcc2b3b181bf43b1ef821ce2",
                    "0x343fdbf42d8aa1136713a119a211d58cec96ef361141d9070a7dfd7e4b9403da"
                ],
                "tx_time": "Jun 17, 11:41 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "a88ead7b-2988-40bb-8d83-59629ba08484",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0477 ETH",
                "val": 0.0477,
                "tx_hash_list": [
                    "0x417b9298617f69b88a9b5016a963f7978d468d756618c1576cd93a131b2d90c7",
                    "0x2c0500e064b17dd1432e69b291e9ba0610d6fa2894ef3063f9d1a50f589c3f0f",
                    "0xffd5d1704b6c875b15321ddf78e973e60c9215cc02ce637e0ab4d8cf6306ba18",
                    "0x130f3ad995d0df3f23040965672dd6b350eb198866b0cc461cdb79d2de7986f9"
                ],
                "tx_time": "Jun 17, 11:11 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "9d229993-ae94-4246-89b5-203e1118c39e",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0008 ETH",
                "val": 0.0008,
                "tx_hash_list": [
                    "0xf54871490a95230d8c8afb362f4774713ce5be5f142266fe16cb1eb305ef202d"
                ],
                "tx_time": "Jun 15, 10:15 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "b3b9f010-32d2-4a1d-94f3-1f054cdd8870",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.006 ETH",
                "val": 0.006,
                "tx_hash_list": [
                    "0x21c1f3863edd361564d5486129fcc8c5fb68a1acd95dc3600b3a5e7ade51a71b",
                    "0x13419800ba44300f990a19ca3762c47db1a6664768accc5822797f3b9303c76d"
                ],
                "tx_time": "May 16, 11:51 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "fced5389-9745-484a-bdd6-32c08bb830e8",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0005 ETH",
                "val": 0.0005,
                "tx_hash_list": [
                    "0xe19b78b226122b353d544bba67102a9b8776d37dadd40a1ff6660369b76b22b2",
                    "0xfe6cfcc43bc9a2ac0acc17d4b6cf02c7e204425fb823716f6e0f72222bdd0d72"
                ],
                "tx_time": "Apr 24, 02:33 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "d6a6a547-2592-4a9a-be7b-e7dacd3b03a4",
                "label": "0.0313 ETH",
                "val": 0.0313,
                "tx_time": "Apr 24, 12:51 AM",
                "tx_hash_list": [
                    "0x71cc3551c61bc5236046ce13493637304a6a728c3f430b0bc5ed8f260f887107",
                    "0x5e541d72707476469ddd9de2b0713566687c0548603b73a89d9519f25b195024",
                    "0xef1ab3dd2dd1abc2297e530a47a7f1a7690abaa9eceabee37943ab11fb6f2a88"
                ]
            },
            {
                "from": "349b465c-7708-41a4-8e05-c7f4019fbc87",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.003 ETH",
                "val": 0.003,
                "tx_hash_list": [
                    "0x10c19307738afb555fc0936e0869d0fded064345aa7a79b034bd79f03ee07430"
                ],
                "tx_time": "Apr 12, 12:49 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "4c5d4d83-e2e3-48f9-b1d1-825e55aa1d16",
                "label": "0.0 ETH",
                "val": 0.0,
                "tx_time": "Mar 30, 12:29 PM",
                "tx_hash_list": [
                    "0xcd66d365cefd01e592df92765b33bcca929ae9060159dcc933ece1a64d9f4c5d"
                ]
            },
            {
                "from": "3762654b-27e7-4dd9-adb6-356ff468051b",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0003 ETH",
                "val": 0.0003,
                "tx_hash_list": [
                    "0x07334727ac45ae49c35e956104e2637b09de9ebbf4d44a47d0cf4ebe116e4e3a"
                ],
                "tx_time": "Mar 30, 12:03 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "6f8b2865-87f0-4a5f-bfa1-5c2818208a40",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0146 ETH",
                "val": 0.0146,
                "tx_hash_list": [
                    "0x5ed2bde475e767410dc946576ff0a09801e0a10fa8efa8fcae89a0bb7875fd1e"
                ],
                "tx_time": "Mar 18, 12:24 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "78436200-4927-4006-91a6-653b2c178621",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.002 ETH",
                "val": 0.002,
                "tx_hash_list": [
                    "0x5b994c7ad10ca5e1371720e14feabb4cadaa56198d08c914df6b1e0a7af430be"
                ],
                "tx_time": "Mar 13, 06:40 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "8435fb7b-7878-42e6-8f37-ad87ea66d6be",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0294 ETH",
                "val": 0.0294,
                "tx_hash_list": [
                    "0xabdbf2b3e3e13cd6a4fdaf2c7091a34f5ded5d948fe3633521b5c36e95cb6823"
                ],
                "tx_time": "Dec 30, 2024, 12:51 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "2af28f8d-a08f-4fce-8f59-f4557cf66e8b",
                "label": "0.0418 ETH",
                "val": 0.0418,
                "tx_time": "Dec 28, 2024, 09:00 AM",
                "tx_hash_list": [
                    "0xab34f590125492c6e497e973ea072427c3dd985074521e0f43d579c3f3da472e",
                    "0x03134de2631c90033c41ef9403f1aad4dffd5502b8cf255c20cce809146c7976",
                    "0x12f624ba1b2218ab0de22495917140d0f918124a60eae08066167296a31df208"
                ]
            },
            {
                "from": "eb6ff87d-c116-40e3-9c30-109cf01698b3",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.04 ETH",
                "val": 0.04,
                "tx_hash_list": [
                    "0xd66987178ffc3a42a62e67f177f3bee96b1d12de3ec81209ebf7aa248f27e093"
                ],
                "tx_time": "Dec 28, 2024, 08:59 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "a25b527c-e911-4ae4-a531-40f3eb6277f3",
                "label": "0.0026 ETH",
                "val": 0.0026,
                "tx_time": "Dec 04, 2024, 03:22 PM",
                "tx_hash_list": [
                    "0x388a2463afa8695a7e57266efdfedcdbf9560256841c55631b7c4187afc2dd7a",
                    "0x9000de87694fb73c792cb278d537099b76288000a17ecc0007f56c0f162afc77"
                ]
            },
            {
                "from": "47843943-cb8e-48ed-a850-655ce7711ca9",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0048 ETH",
                "val": 0.0048,
                "tx_hash_list": [
                    "0x1fe6917beae061d116a065fe5f73863f81fabbd6656449ffab04edb66d2b1e50"
                ],
                "tx_time": "Dec 04, 2024, 03:19 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2a180f1c-78d9-4dc0-b30f-36b354c53787",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0012 ETH",
                "val": 0.0012,
                "tx_hash_list": [
                    "0x67ab7817a5b170f4d90e8f900379ba71faed950fb16e91a6af5ee0e31776b482",
                    "0xd571c61eb573385e0a4d5a5feb30fbc938c5ef3f19d9e57cb93eed86bd2e7775"
                ],
                "tx_time": "Nov 09, 2024, 01:57 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "466833c3-5332-40ff-9230-4c484f9bc1ca",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.003 ETH",
                "val": 0.003,
                "tx_hash_list": [
                    "0x25e7e1bdcd0070caf7d831bf97fb6a8eba1f4d2101ff69706d3cc6ca9e8eb869"
                ],
                "tx_time": "Nov 06, 2024, 02:27 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "a31807b0-710e-4256-a40a-25f35f602019",
                "label": "0.0001 ETH",
                "val": 0.0001,
                "tx_time": "Oct 12, 2024, 04:35 AM",
                "tx_hash_list": [
                    "0xff296050ccae8c9686e0f9c3595d1afc4851ceab76ded766b76745cec4ed0873"
                ]
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "9853d89d-85f4-41df-9e2f-bf0dbcad3d9c",
                "label": "0.027 ETH",
                "val": 0.027,
                "tx_time": "Oct 12, 2024, 03:37 AM",
                "tx_hash_list": [
                    "0xefb9f31b1b2c040f317a05085f581849c36ea0ddf14a03ea8ecce8902e450c21"
                ]
            },
            {
                "from": "b1e14a67-fdc1-4103-9ca5-ad9ba979e561",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0283 ETH",
                "val": 0.0283,
                "tx_hash_list": [
                    "0xd89bc473dcc894e1377ada15d878b7c9dcd3b862bacbb3e35a1b128f1de35feb"
                ],
                "tx_time": "Oct 11, 2024, 10:49 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "543a26b2-e206-4794-9cc0-45761c165a9a",
                "label": "0.0317 ETH",
                "val": 0.0317,
                "tx_time": "Aug 31, 2024, 04:37 AM",
                "tx_hash_list": [
                    "0xf5d0f2c78f83acf827d33004132d8330619d7880a263017dc4569f4771985f20",
                    "0xe3263ef24fb3a4f8b7fdeb6c208c09e768ca5e3d01082ff28738d8a9c985e243",
                    "0x2abf0cf8183742ea3578791a57fd82aa2a412b4d899b1a69eb5b00a3e8bba1d5",
                    "0x1570c6fd9020c431666cb5f7a10f237b53fafbafbf1293a7a4a893545347c506",
                    "0xfd823756b69862095789f4e0ef25218bda6fb865f510956abab62816bf1dc6ac",
                    "0x1ad9328e170d19dfbf12a2bbe0f58b4d06b79d3c19f24b32d035788788d320b8"
                ]
            },
            {
                "from": "aab7135e-b03b-401d-bafc-fb716aff0bff",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.004 ETH",
                "val": 0.004,
                "tx_hash_list": [
                    "0x7c9a25044815f56811405551d4ddb98f0d2cc9636f424b217d30ee024f8f84c2"
                ],
                "tx_time": "Aug 31, 2024, 04:37 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "fcbd92f2-3cf4-485e-aa9b-cddf74779ac8",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.004 ETH",
                "val": 0.004,
                "tx_hash_list": [
                    "0xadedda53c34597d4451957e8825753b55b033dab96cf8643c6773431a022120b"
                ],
                "tx_time": "Aug 20, 2024, 01:48 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "94615aed-467a-4105-9925-882e6d281fa9",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0243 ETH",
                "val": 0.0243,
                "tx_hash_list": [
                    "0x5ec8590ce95b885163f76b6741457846944148666754ce93f0b3e2c1923854e0"
                ],
                "tx_time": "Aug 12, 2024, 12:00 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "b3d4caf6-4904-4930-8ac9-a042a9449bc4",
                "label": "72.7015 ETH",
                "val": 72.7015,
                "tx_time": "Aug 05, 2024, 07:40 PM",
                "tx_hash_list": [
                    "0xc87424d0c7ea6410f65e1f76f386a39a00c5159719d281093e7629a5ff5c1219",
                    "0xb472d556655570454eccd8bd84bc9b6784a23ecd6de721eb938b9f9ddad73237",
                    "0xde718c29c51a463403adf7794568d3eaecbfd0f9c29ddcbe561e7230b55f6f1b",
                    "0xec72d88b1ebbf227fc10e3be2e2624ed02b73c6687ccec52a47dc3a0910ae91d",
                    "0x72a8bd0e42c01fa117fedfbdd436cf0741d4dfb37fe5fe241dfa5c67a4127daa"
                ]
            },
            {
                "from": "1842f98c-47f4-4409-a756-66277dd3a0c1",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0006 ETH",
                "val": 0.0006,
                "tx_hash_list": [
                    "0x0a1420728da8a35838697c50f58d25f541e11204d7d5702bb60c6156dda11be3"
                ],
                "tx_time": "Aug 05, 2024, 07:40 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "b9e1b93d-808b-40bc-911e-e1e35aa64ab2",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0003 ETH",
                "val": 0.0003,
                "tx_hash_list": [
                    "0x0839eda22ad7c4ff97b31083ea509d5d1f16a95ccc98ebc1e0c432d466b00946"
                ],
                "tx_time": "Aug 03, 2024, 09:29 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "c6d3e28d-768d-4442-b47c-492e26154c83",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.01 ETH",
                "val": 0.01,
                "tx_hash_list": [
                    "0xe1f8ee389e60ae97d0d759366a1e3639ae515525715b5107878a3f21a8719072"
                ],
                "tx_time": "Jun 17, 2024, 06:21 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "7af9790b-eda7-4efe-8ab7-a542b0ff4d40",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.001 ETH",
                "val": 0.001,
                "tx_hash_list": [
                    "0xbe4e9c03a3bd2cac32079327ccf8c82064bcb4540b5094c052bf186d614cbb6b"
                ],
                "tx_time": "Jun 17, 2024, 08:36 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "84109363-dc7b-4fd7-aaa0-57c231541121",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.002 ETH",
                "val": 0.002,
                "tx_hash_list": [
                    "0xa4c304b1d671d91872f3e866d3578481258ab3a5a2df5726ab1863a4b2b2c343"
                ],
                "tx_time": "May 30, 2024, 01:56 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "bed23184-5c45-470e-83dc-6144d2393b9e",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.002 ETH",
                "val": 0.002,
                "tx_hash_list": [
                    "0xa5bc17181ae8ae0708a96f17ba8343f3f2525297683ec26b83c79706655ec926",
                    "0x40ec86d83f0bea36574e9ae208a133d20e7e9b1d2b86cbbc7f3ec4567338c455"
                ],
                "tx_time": "May 26, 2024, 03:52 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "d4ba3b04-d36d-4178-aa6b-c6ed0a1be828",
                "label": "0.0006 ETH",
                "val": 0.0006,
                "tx_time": "May 23, 2024, 02:09 PM",
                "tx_hash_list": [
                    "0x4afdcd73051b4b3c3c3232598733854c8be5535c58194c59356112f3a840648b"
                ]
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "0912a14f-f011-4c82-b769-e044ae0441cf",
                "label": "0.39 ETH",
                "val": 0.39,
                "tx_time": "May 10, 2024, 07:59 PM",
                "tx_hash_list": [
                    "0x789a075f8da3bef76583024919abb26701248a56fa243a833f54e796eba3cd72"
                ]
            },
            {
                "from": "fa000a3e-e24b-4cbf-afba-4f00de45963f",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.4 ETH",
                "val": 0.4,
                "tx_hash_list": [
                    "0x2cc8c7737c8882b13be7525c70dcb6cddbc60df1265d565fec885beb6714c1ac"
                ],
                "tx_time": "May 10, 2024, 06:39 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "14f07197-0ebb-4b12-8499-000492964874",
                "label": "0.0755 ETH",
                "val": 0.0755,
                "tx_time": "Apr 05, 2024, 07:38 AM",
                "tx_hash_list": [
                    "0x594fe796559d3dcef1cf66796c7572d473f71a66040ace5f180f4ab44554ffcf",
                    "0xfbe0842c9fa69581c76e7f9363d359c214f37944f16089f96ff744f03131663e",
                    "0x469562fb363f8267ceeb908dcd204e1a4de1546e5c5f144529e5b79edd33860e",
                    "0x218687867662bc7ea5c9d7af175b011a7fd1879733ff2d830ed21a92162a42bb",
                    "0x0572b00361eea677a6caa1b56f2f1c5dde2454e096916b6287537582ee6c63e9",
                    "0x7c5bfacb289fd6d835370d33215b97f0299dba531a2808f54542f9d76ebae72b",
                    "0x014362011a770f59855f2cb2576f86159390e381b7e4c14ba6d5d07bc057c97c",
                    "0x6cee2a5080439e5a95e39f0835f86275e8501f5ea00b41a6f190ec6922130d9b",
                    "0x6691cda072ce0c16236462ffb7b2b1ecef22787d2903a473a7f798193fb41bf9",
                    "0xca302cf77403bf6177ced2834248a08c754dae7482961d6a65d68fbb7201cd65",
                    "0x1fea5c65ef0186c9e2d59e571d1d7bfd79ef821b5e69f60395c1c091c187f5bc"
                ]
            },
            {
                "from": "a61d54fb-8890-4f8d-a487-bcb401de43cd",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.01 ETH",
                "val": 0.01,
                "tx_hash_list": [
                    "0x75c93ed97f707b34dfb1b03d640bacc1b3f46628773fe3206ee06a0d2dbe4005"
                ],
                "tx_time": "Apr 05, 2024, 07:37 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "dbecbba4-00ff-4053-b8b4-f20737a5c65d",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.1 ETH",
                "val": 0.1,
                "tx_hash_list": [
                    "0xf72f061941fd10d9a67dcb52566e421eddbd3ef0789836ac52155c6fd0ba82b6"
                ],
                "tx_time": "Mar 31, 2024, 01:11 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "a3b0feef-4c98-42e9-bb0f-19f1054b9648",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0036 ETH",
                "val": 0.0036,
                "tx_hash_list": [
                    "0x89ec9dc580686ce62b63c5655c04a52ab9f997d06478c5cd22a1453c1e7fc4ff",
                    "0x8ebdb85090ad2fdef1ed517e707aa1d59fc506461ded3606c0063390be019e1c"
                ],
                "tx_time": "Mar 30, 2024, 12:48 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "a720f9a9-ae5e-4cb6-96ce-b7d621b7e4bc",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.013 ETH",
                "val": 0.013,
                "tx_hash_list": [
                    "0x5356d14fea795abca8dbdc822cd6205371297544e6bc0f737657ee904c53c9ac"
                ],
                "tx_time": "Mar 30, 2024, 11:58 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "b4f90922-170c-4034-b33a-11da69144049",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0109 ETH",
                "val": 0.0109,
                "tx_hash_list": [
                    "0x6528959fadaf07dc88869388224144a0616328e9dc5785c84a13e793b7c9f808"
                ],
                "tx_time": "Mar 28, 2024, 01:36 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "c793fd3a-63b6-4996-82d9-070882052f56",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.055 ETH",
                "val": 0.055,
                "tx_hash_list": [
                    "0xcb7e99be52faeebe2d1e7bf3fbcc669bc2861f10c70b103a828d9eadcedc1490",
                    "0xc5d93d0b3e64e333bf5071fa5634a5d5749f6243898b235f339d8e2ad0789739",
                    "0xa8fb86df6714cd5e82174653adcd7178cde56f01c9d431dc55a175d83d40adbd",
                    "0x34a30af55495a28fe30db95b6d3c6119fa3c8b8a0f66f7b66195daca472f8044",
                    "0x6ae28a801c522f36c2a39f62a33d9eced0807595a12066537d104f015dbd32a3"
                ],
                "tx_time": "Mar 27, 2024, 11:10 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "eb86cc20-5e7d-4bba-9f98-0cf4546590e1",
                "label": "0.0357 ETH",
                "val": 0.0357,
                "tx_time": "Mar 25, 2024, 08:17 AM",
                "tx_hash_list": [
                    "0xefb5c233961b3b30a8d6541ea8814345a4a3cc5aa7715d20096dcdb45e4b63bd",
                    "0x2180f1741c6bd0f29b3cad582250ffb9ba0cca418192c1a1554dca80afe33d93",
                    "0xe06f694d64d432b9274cb3d61688a0226c4b46123818c8301913728fa0fedcfe",
                    "0x7ba946414a5275e62bda270e4742af271fdad4414fb1dab1a5a399886862d45a"
                ]
            },
            {
                "from": "cef9001b-8449-4db3-a6c9-19d108d5bff7",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.001 ETH",
                "val": 0.001,
                "tx_hash_list": [
                    "0xd4ffdef99ff11180844e0bf16ae4ec750504275efeb92270f24ff8ac7a6007ca"
                ],
                "tx_time": "Mar 25, 2024, 08:17 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "c2d7232b-d64d-4706-9174-77e5287948fd",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0162 ETH",
                "val": 0.0162,
                "tx_hash_list": [
                    "0x099c3b0f34efd1d1cd53fe8780cdcb24c95a250cbcfa45d74e1e3a00997cae69"
                ],
                "tx_time": "Mar 25, 2024, 03:07 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2e6e7fd6-f066-4ea9-ae23-539a84d3fe8b",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0018 ETH",
                "val": 0.0018,
                "tx_hash_list": [
                    "0xcc0b955ea133ba85687b08dca1df75161e373fa20c7bdadde0c7c1cd14e2a5d9"
                ],
                "tx_time": "Mar 24, 2024, 08:06 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "40bd1521-83bb-445b-b124-a0e5dada0d4f",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0056 ETH",
                "val": 0.0056,
                "tx_hash_list": [
                    "0xe2d8f678a4171b00fde2705b7d6ae2ff42d23b68eecef92f0ccb25acb3a56800",
                    "0xa70bdefa3c374331409522bd864b53c0f3a8e22b456d66aa2b86fc1bc7b2a89b"
                ],
                "tx_time": "Mar 24, 2024, 07:32 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "5809ed55-956d-48b7-92c1-834799eaaf3d",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.002 ETH",
                "val": 0.002,
                "tx_hash_list": [
                    "0x0d86ad84e7bec7484a3f9594b1eeaab7f6b9c87d86b8a16e5318c34db10fb8df"
                ],
                "tx_time": "Mar 24, 2024, 07:08 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "d9d6dcc5-ac3f-49ff-b0f6-b7ad57afb5d2",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.02 ETH",
                "val": 0.02,
                "tx_hash_list": [
                    "0xb279bad3cae16577adc9dd3418858ead611aa000ce56dbc4b9ef0733bdc19ecd"
                ],
                "tx_time": "Mar 19, 2024, 02:12 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "6282668b-b44a-452b-8e01-59c0f9ed618b",
                "label": "0.0095 ETH",
                "val": 0.0095,
                "tx_time": "Mar 18, 2024, 04:32 AM",
                "tx_hash_list": [
                    "0xe7a408152a8742e5552f04000562cb09fc0578a6b5c622542e19b4be84e5c558"
                ]
            },
            {
                "from": "c9b2a0a0-4b90-44af-bf55-7d27614d4d09",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.205 ETH",
                "val": 0.205,
                "tx_hash_list": [
                    "0x7e8982dcff63d21da86ae592b9164c97e77dc713ca81fad46f5585e9e14f7a24",
                    "0xb79ee3f6849409101421f66fa357e32cb97b9a07cc4373e4a6005df6e84e362e"
                ],
                "tx_time": "Feb 19, 2024, 05:06 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "a286cea3-8a1d-4a2c-bf12-a0cfeb0034ac",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.01 ETH",
                "val": 0.01,
                "tx_hash_list": [
                    "0x10784cb2b6044a37bc05663455b0597d0f5abbd409bf2ef3290181cc25f86827"
                ],
                "tx_time": "Feb 04, 2024, 12:50 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "c6a638ba-5c20-44c0-9a99-22add03bb1ff",
                "label": "0.028 ETH",
                "val": 0.028,
                "tx_time": "Feb 02, 2024, 11:42 PM",
                "tx_hash_list": [
                    "0x0e67ccbb5c2d20541e891436c8e2557e4ea7eea21c6db86aa315ebabb3dc5297",
                    "0x76429cc6f799ac28abe7ab0c8d5822b1236d1b58a03010a3d7db1a7812f2a642",
                    "0x866fa40ce9dabd401cf1c81cd1f705aee3511881f2c0057ea0a773b35905248a",
                    "0x482a5d2a2f12022754c10cdca8ec058bc29c22fc0fe4c6f2afbd7fef37ee03b0"
                ]
            },
            {
                "from": "8b318e8d-4481-4259-8b6c-841b6f93619d",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.013 ETH",
                "val": 0.013,
                "tx_hash_list": [
                    "0x381f65ffad108c19f45efdfe6b3b4e37bb9479c0d8e2262adb576db06f2f7429"
                ],
                "tx_time": "Feb 02, 2024, 11:39 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "ffcba112-b1be-426c-972d-c60b8de6801c",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.008 ETH",
                "val": 0.008,
                "tx_hash_list": [
                    "0xab3a7dce5d57bb231420fb52ac8f1be5b93bbce85fd1ff2af2775ac31ac85b33",
                    "0xf01fb2ada189fa1efe99729053b4efc13fdd7205ab69357c1e58eb7d570ed284"
                ],
                "tx_time": "Feb 02, 2024, 05:40 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "b61d25a1-ec0a-4fd2-83ec-80ccd9cb9047",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.01 ETH",
                "val": 0.01,
                "tx_hash_list": [
                    "0x550e9be753efd9d3e256889ac3c77a28f088f09ebb2758c865485cc052e2d357"
                ],
                "tx_time": "Feb 01, 2024, 07:55 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "ad8628e9-8640-4759-b974-83f5f6290aa4",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0044 ETH",
                "val": 0.0044,
                "tx_hash_list": [
                    "0x7977d93da55462e98b6d322a6c932a63ced49bc6d304118ffce386f1e3e896ef"
                ],
                "tx_time": "Jan 29, 2024, 01:50 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "4c8ed30e-216d-4831-be07-835049d34c59",
                "label": "0.01 ETH",
                "val": 0.01,
                "tx_time": "Jan 29, 2024, 04:49 AM",
                "tx_hash_list": [
                    "0x778b4d3618f8bedaac469eefc8eb802fed43ea1d67bc8fe0009d94c2c49037ea"
                ]
            },
            {
                "from": "e401dafa-19b6-4cae-a1cc-b98acaf18d6d",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.01 ETH",
                "val": 0.01,
                "tx_hash_list": [
                    "0x73a7996cde588dd6fdf8ddab654ba1c664f23261b1cffc867c58c93593bc8212"
                ],
                "tx_time": "Jan 29, 2024, 01:07 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "0f3e5735-f82f-4858-827b-6b0cb8fa84e7",
                "label": "0.01 ETH",
                "val": 0.01,
                "tx_time": "Jan 22, 2024, 04:23 PM",
                "tx_hash_list": [
                    "0x03ef21b06ad9cd5ee3ae33f2e9fcd7aba31ea63bd6c7422ff0af1814edc2d6e5"
                ]
            },
            {
                "from": "06326d94-a1a6-4233-a0dc-0bdc68319ec6",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0125 ETH",
                "val": 0.0125,
                "tx_hash_list": [
                    "0x3622435e48f5a94001f47d6ef80734caeeb7b445b42353235f14d7f5a44d5e29"
                ],
                "tx_time": "Jan 22, 2024, 04:10 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "31e47172-12ce-4729-90f7-5dfdb44bada6",
                "label": "0.2992 ETH",
                "val": 0.2992,
                "tx_time": "Jan 18, 2024, 04:30 AM",
                "tx_hash_list": [
                    "0x3dd584164205fa6bc19202ee8dd7bbd22f3ceb44648876b069907c741624ec0e",
                    "0x532be621e17ea4babee1e420f24d83c89e0b863a5b31f87161c42c37837c64c1"
                ]
            },
            {
                "from": "2ad7218f-296f-4c10-aa61-f87c4a30ac25",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.1 ETH",
                "val": 0.1,
                "tx_hash_list": [
                    "0xdd69424eb5ffe6c939885bccfefe5720397373899547430c9c4c6d0ecd3246d3"
                ],
                "tx_time": "Jan 16, 2024, 04:31 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "e8bbf4e7-e0a5-4f16-bc76-ef461c082a59",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.006 ETH",
                "val": 0.006,
                "tx_hash_list": [
                    "0x88d54d0fdf48f6e18d0d283b599308394ea88cd43ca0032b73ef84948158a11c"
                ],
                "tx_time": "Dec 29, 2023, 07:28 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "f137ffd5-b153-420e-87a8-edc69b39d751",
                "label": "0.0299 ETH",
                "val": 0.0299,
                "tx_time": "Nov 22, 2023, 12:57 AM",
                "tx_hash_list": [
                    "0xa52017a02534ccf9a40cd3b60c9cb40642231c657a0ebf592e18494628487b77",
                    "0x0392dba447c8eeaff027786d8542b2137d23ac1c99f76492d57acfc4d885cd41",
                    "0xa8dc667121e261474443930a90fa334397b02cd8a44a4e56e8e63c391d9eeed2",
                    "0x51172f741441f9eec72076903319d31d8941437520167d34f2919a981731105e",
                    "0x0fbc3039243efbad1bd4d66cdb01835c563e57a37d7f541c17e6b896eec13f26"
                ]
            },
            {
                "from": "61cc42a0-285e-421b-82e7-c699cf9adf2a",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.004 ETH",
                "val": 0.004,
                "tx_hash_list": [
                    "0xad68d7a0b5466771517479db85caeebf78aedcc4d7bc398b7cd6fa220fbf7f5a"
                ],
                "tx_time": "Nov 22, 2023, 12:56 AM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "bcfcbda8-114d-4911-97b0-bae644b10805",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.018 ETH",
                "val": 0.018,
                "tx_hash_list": [
                    "0x9cfe0f33ded05cbb2f773a237f2929846cc508941d5cb45d3074142b4545d8b7",
                    "0x254744c7e21b5612d284e2a565e034a55ddd98f27b6d4d83dab81858d6a2b047",
                    "0x2e9ecb79c9457069ba99367f94627a2211001d1bc7d0a15a2b85ff382941b998"
                ],
                "tx_time": "Nov 21, 2023, 04:38 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "56b41b43-8058-49c2-80a7-ad3d25e51ad4",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.012 ETH",
                "val": 0.012,
                "tx_hash_list": [
                    "0x843403b474f92b5b6241e638bf064c634a78e3d2dcd10af5c04cc338a6edd6ad",
                    "0x11097684e268fa3b93a94c68a72ed5f91a43280e03442de1e010727f0f9bd41b"
                ],
                "tx_time": "Nov 21, 2023, 04:32 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "100150cf-61ff-4b0c-b328-480730e8301f",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.002 ETH",
                "val": 0.002,
                "tx_hash_list": [
                    "0x61cd58968c95c205b13fbc6cf0e79a2dd6ff22d612060107773a5aa1724f68b4"
                ],
                "tx_time": "Nov 21, 2023, 04:21 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "03229d72-5ad1-4815-9bbc-19d23fe49cbe",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "0.0026 ETH",
                "val": 0.0026,
                "tx_hash_list": [
                    "0x106c7d9dc8ed473888b6958401c2409a6202b361ffe4d51b8113ec99a5b92bb4"
                ],
                "tx_time": "Nov 21, 2023, 04:04 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "826e97d9-fdbd-4100-b8e3-a34e65937f06",
                "label": "0.0212 ETH",
                "val": 0.0212,
                "tx_time": "Nov 18, 2023, 10:56 PM",
                "tx_hash_list": [
                    "0x2b5564ebd539c0c26e806cf63438832393ad7230d0499c44691d9e4a1b28733d",
                    "0xf01b835a16e98cb74ed6b3152280b49d89f20059d0bc3057518bd62d9145a293",
                    "0xce77cb5626efdea3100c63e8a6fbdb61aaa52aa6548ba681e468318466d774f3"
                ]
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "2288997a-6ca3-4a13-a1f4-282a1ff7bf49",
                "label": "0.02 ETH",
                "val": 0.02,
                "tx_time": "Aug 19, 2023, 10:49 AM",
                "tx_hash_list": [
                    "0xb7842d3a74d1018651ac2f6e9027c9829e3b650bd3a51d1e6fae3645233b9bb3"
                ]
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "6e4ee2d5-9aca-4f20-8e81-0cda0c57900e",
                "label": "0.005 ETH",
                "val": 0.005,
                "tx_time": "Aug 19, 2023, 05:56 AM",
                "tx_hash_list": [
                    "0x88233387e12f9844e1cf6399a77a30e3595249b665d16fc5aa8a047058a0675e",
                    "0x8c6652e86f042c0703b4d5338cc2f2d9d4a570e63fca517343eab92906e202c7"
                ]
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "dc15b9cb-0f42-42e8-ad3c-6cd4c523eae7",
                "label": "0.01 ETH",
                "val": 0.01,
                "tx_time": "Aug 19, 2023, 05:48 AM",
                "tx_hash_list": [
                    "0xf4f59be2dc56424cba6dd8aedf98e7d9137fa7be25961c93981bea90e52f59b1"
                ]
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "21898669-9a6d-4d7a-b49c-1a1b45798ae5",
                "label": "0.08 ETH",
                "val": 0.08,
                "tx_time": "Jul 24, 2023, 03:08 AM",
                "tx_hash_list": [
                    "0xc53a8b17fecb546d3fcbc1783999f12c778094c0f2de4bc2e7dda1acbeba1d2e"
                ]
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "2240c595-f899-41c0-afc0-c7e0556c5fe8",
                "label": "1.0 ETH",
                "val": 1.0,
                "tx_time": "May 26, 2023, 10:47 AM",
                "tx_hash_list": [
                    "0xcbc6417e9d18154bccbf87eb7c08bbfa10654806cf6ccf167867dacdad3562ad"
                ]
            },
            {
                "from": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "to": "e729e355-cc6b-429c-bcd6-d4373baa3b0f",
                "label": "1.0 ETH",
                "val": 1.0,
                "tx_time": "May 25, 2023, 06:34 PM",
                "tx_hash_list": [
                    "0xc592e64fb32f100bf16d0f160e9a7b6531429fbe7b58412078dacf67d0528706"
                ]
            },
            {
                "from": "d1936264-6221-45f4-b1f4-2bbf5adf74e4",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "4.522 ETH",
                "val": 4.522,
                "tx_hash_list": [
                    "0xa3907acbfaf297a91003f8716c46c25c950fb41622fd648ddfdaf13dbead6231"
                ],
                "tx_time": "Jul 01, 2022, 09:10 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            },
            {
                "from": "c18e8de7-b5c9-49b9-a92f-e7d3693ac6da",
                "to": "2263909c-9b8b-4b58-ab8e-c891e0e609ee",
                "label": "70.3478 ETH",
                "val": 70.3478,
                "tx_hash_list": [
                    "0x7dfa03bfdc4e91e10fb5f147c3fafa2f9592148acd836f9990ff4ca2cf43673a"
                ],
                "tx_time": "Jul 01, 2022, 09:03 PM",
                "color": {
                    "color": "#99b4d7",
                    "highlight": "#99b4d7"
                }
            }
        ],
        "tx_count": 210,
        "first_tx_datetime": "2022-07-02",
        "latest_tx_datetime": "2025-08-09",
        "address_first_tx_datetime": "2022-07-02",
        "address_latest_tx_datetime": "2025-08-09"
    },
    "address_first_tx_datetime": "2022-07-02",
    "address_latest_tx_datetime": "2025-08-09"
}`)

	var labelAddresses LableAddresList
	err := json.Unmarshal(jsonData, &labelAddresses)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}

	if len(labelAddresses.GraphDic.NodeList) > 0 {
		for _, data := range labelAddresses.GraphDic.NodeList {
			fmt.Println("data label: ", data.Label)
			if strings.Contains(data.Label, "huione") {
				fmt.Println("汇旺")
			}
			if strings.Contains(data.Label, "Theft") {
				fmt.Println("盗窃")
			}
			if strings.Contains(data.Label, "Drainer") {
				fmt.Println("=诈骗")
			}
			if strings.Contains(data.Label, "Banned") {
				fmt.Println("制裁")
			}
		}
	}
}
