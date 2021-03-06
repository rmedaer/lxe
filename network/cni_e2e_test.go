package network

// func testCNIPlugin(t *testing.T) *cniPlugin {
// 	t.Log("testCNIDir", testCNIDir)

// 	cniPlugin, err := InitPluginCNI(ConfCNI{testCNIbinPath, testCNIconfPath, testCNInetnsPath})
// 	assert.NoError(t, err)

// 	return cniPlugin
// }

// func toFakeCniPlugin(_ *testing.T, cniPlugin *cniPlugin) *libcnifake.FakeCNI {
// 	fake := &libcnifake.FakeCNI{}
// 	cniPlugin.cni = fake

// 	return fake
// }

// func Test_cniPlugin_PodNetwork(t *testing.T) {
// 	cniPlugin := testCNIPlugin(t)
// 	type args struct {
// 		id          string
// 		annotations map[string]string
// 	}

// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{"", args{"foo", nil}, false},
// 	}

// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := cniPlugin.PodNetwork(tt.args.id, tt.args.annotations)
// 			pn := got.(*cniPodNetwork)
// 			fmt.Print(err)
// 			assert.False(t, (err != nil) != tt.wantErr)
// 			assert.NotNil(t, pn.netList)
// 			assert.NotNil(t, pn.runtimeConf)
// 		})
// 	}
// }

// func Test_cniPodNetwork_Attach(t *testing.T) {
// 	testNeedsRoot(t)
// 	cniPlugin := testCNIPlugin(t)
// 	podNetwork, err := cniPlugin.PodNetwork("test_attach", nil)
// 	assert.NoError(t, err)
// 	containerNetwork, err := podNetwork.ContainerNetwork("containerid", nil)
// 	assert.NoError(t, err)
// 	fake := toFakeCniPlugin(t, cniPlugin)

// 	_, err = containerNetwork.WhenStarted(ctx, &PropertiesRunning{Pid: 0})
// 	assert.NoError(t, err)
// 	assert.Equal(t, 1, fake.AddNetworkListCallCount())

// 	out, err := exec.Command("ip", "netns", "delete", "test_attach").CombinedOutput()
// 	assert.NoError(t, err, string(out))
// }

// func Test_cniPodNetwork_Teardown_MissingNetwork(t *testing.T) {
// 	testNeedsRoot(t)
// 	cniPlugin := testCNIPlugin(t)
// 	podNetwork, err := cniPlugin.PodNetwork("test_teardown_missingnetwork", nil)
// 	assert.NoError(t, err)
// 	containerNetwork, err := podNetwork.ContainerNetwork("containerid", nil)
// 	assert.NoError(t, err)
// 	//fake := toFakeCniPlugin(t, cniPlugin)

// 	// CRI DelNetwork always tries to remove as good as possible without throwing error
// 	err = containerNetwork.WhenDeleted(ctx, nil)
// 	assert.NoError(t, err)
// 	//assert.Equal(t, 1, fake.DelNetworkListCallCount())
// } // nolint: wsl

// func Test_cniPodNetwork_Status_MissingNetwork(t *testing.T) {
// 	testNeedsRoot(t)
// 	cniPlugin := testCNIPlugin(t)
// 	podNetwork, err := cniPlugin.PodNetwork("test_status_missingnetwork", nil)
// 	assert.NoError(t, err)
// 	//fake := toFakeCniPlugin(t, cniPlugin)

// 	got, err := podNetwork.Status(ctx, &PropertiesRunning{})
// 	assert.Error(t, err)
// 	assert.Nil(t, got)
// 	//assert.Equal(t, 1, fake.CheckNetworkListCallCount())
// } // nolint: wsl

// func Test_cniPodNetwork_Status_WithNetwork(t *testing.T) {
// 	testNeedsRoot(t)
// 	cniPlugin := testCNIPlugin(t)
// 	podNetwork, err := cniPlugin.PodNetwork("test_status_withnetwork", nil)
// 	assert.NoError(t, err)
// 	containerNetwork, err := podNetwork.ContainerNetwork("containerid", nil)
// 	assert.NoError(t, err)
// 	//fake := toFakeCniPlugin(t, cniPlugin)

// 	_, err = containerNetwork.WhenStarted(ctx, &PropertiesRunning{Pid: 0})
// 	assert.NoError(t, err)
// 	//assert.Equal(t, 1, fake.AddNetworkListCallCount())

// 	got, err := podNetwork.Status(ctx, &PropertiesRunning{Properties: Properties{Data: map[string]string{"result": `{"cniVersion":"0.4.0","ips":[{"version":"4","interface":2,"address":"10.22.0.64/16","gateway":"10.22.0.1"}]}`}}})
// 	assert.NoError(t, err)
// 	assert.NotNil(t, got)
// 	//assert.Equal(t, 1, fake.CheckNetworkListCallCount())

// 	out, err := exec.Command("ip", "netns", "delete", "test_status_withnetwork").CombinedOutput()
// 	assert.NoError(t, err, string(out))
// }
