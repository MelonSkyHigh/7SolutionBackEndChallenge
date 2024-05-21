package MeatCounter

import "testing"

func TestMeatFormat(t *testing.T) {
	type args struct {
		baconipsum string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test_01_Normal",
			args: args{"Pork loin anim veniam, landjaeger dolore laboris shankle voluptate.  Shank landjaeger porchetta andouille swine et mollit est cow quis tri-tip officia.  Ipsum fugiat dolore enim incididunt, in cow short loin tongue short ribs qui labore beef.  Incididunt bresaola shank commodo hamburger sint jerky dolore sed salami chislic ground round deserunt ball tip."},
			want: "{\"andouille\":1,\"anim\":1,\"ball\":1,\"beef\":1,\"bresaola\":1,\"chislic\":1,\"commodo\":1,\"cow\":2,\"deserunt\":1,\"dolore\":3,\"enim\":1,\"est\":1,\"et\":1,\"fugiat\":1,\"ground\":1,\"hamburger\":1,\"in\":1,\"incididunt\":2,\"ipsum\":1,\"jerky\":1,\"labore\":1,\"laboris\":1,\"landjaeger\":2,\"loin\":2,\"mollit\":1,\"officia\":1,\"porchetta\":1,\"pork\":1,\"qui\":1,\"quis\":1,\"ribs\":1,\"round\":1,\"salami\":1,\"sed\":1,\"shank\":2,\"shankle\":1,\"short\":2,\"sint\":1,\"swine\":1,\"tip\":1,\"tongue\":1,\"tri-tip\":1,\"veniam\":1,\"voluptate\":1}",
		},
		{
			name: "Test_02_UpLowCase",
			args: args{"Pork. pork, PorK...pORK pork PORK chicken"},
			want: "{\"chicken\":1,\"pork\":6}",
		},
		{
			name: "Test_03_Empty",
			args: args{""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MeatFormat(tt.args.baconipsum); got != tt.want {
				t.Errorf("MeatFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
