package sound

import (
	"encoding/json"
	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
	"github.com/xiusin/castr-tickeys/helper"
	"github.com/xiusin/logger"
	"io/ioutil"
	"path/filepath"
)

type ConfItem struct {
	Name           string         `json:"name"`
	DisplayName    string         `json:"display_name"`
	Files          []string       `json:"files"`
	NonUniqueCount int            `json:"non_unique_count"`
	KeyAutoMap     map[string]int `json:"key_auto_map"`
}

var schemes []ConfItem

var streamers struct {
	Ctx       *oto.Context
	Sounds    [][]byte
	Modifiers map[string][]byte
}

var sound ConfItem

func init() {
	_ = json.Unmarshal(confJson(), &schemes)
	sound = schemes[0]
	for _, v := range schemes {
		if helper.GetConf().SoundType == v.Name {
			sound = v
			break
		}
	}
}

func GetSoundConf() ConfItem {
	return sound
}

func InitStreamer() struct {
	Ctx       *oto.Context
	Sounds    [][]byte
	Modifiers map[string][]byte
} {
	soundDir := helper.AppDirPath("../Resources/sounds/" + sound.Name)
	//soundDir := "/Users/xiusin/projects/src/github.com/xiusin/castr-tickeys/sounds/" + sound.Name
	var sounds [][]byte
	var modifiers = map[string][]byte{}
	for idx, soundFile := range sound.Files {
		if idx < sound.NonUniqueCount {
			f := filepath.Join(soundDir, soundFile)
			var file, err = ioutil.ReadFile(f)
			if err != nil {
				logger.Error("打开声音", f, "失败", err)
			}
			_, data, _ := minimp3.DecodeFull(file)
			sounds = append(sounds, data)
		}
	}
	for k, v := range sound.KeyAutoMap {
		var file, _ = ioutil.ReadFile(filepath.Join(soundDir, sound.Files[v]))
		_, data, _ := minimp3.DecodeFull(file)
		modifiers[k] = data
	}
	streamers.Ctx, _ = oto.NewContext(44100, 2, 2, 1024)
	streamers.Sounds = sounds
	streamers.Modifiers = modifiers
	return streamers
}

func confJson() []byte {
	return []byte(`[
	{
		"name":"bubble",
		"display_name":"bubble",
		"files":["1.mp3","2.mp3","3.mp3","4.mp3","5.mp3","6.mp3","7.mp3","8.mp3", "enter.mp3"],
		"non_unique_count":8,
		"key_audio_map":{"36":8}
	},

	{
		"name":"typewriter",
		"display_name":"typewriter",
		"files":["key-new-05.mp3","key-new-04.mp3","key-new-03.mp3","key-new-02.mp3","key-new-01.mp3","space-new.mp3","scrollUp.mp3","scrollDown.mp3","backspace.mp3", "return-new.mp3"],
		"non_unique_count":5,
		"key_audio_map":{"36":9, "49":5, "51":8, "116":6, "121": 7}
	},

	{
		"name": "mechanical",
		"display_name":"mechanical",
		"files":["1.mp3", "2.mp3", "3.mp3", "4.mp3", "5.mp3"],
		"non_unique_count":4,
		"key_audio_map":{"36":4}
	},

	{
		"name": "sword",
		"display_name": "sword",
		"files": ["1.mp3", "2.mp3", "3.mp3", "4.mp3", "5.mp3", "6.mp3", "back.mp3", "enter.mp3", "space.mp3"],
		"non_unique_count": 6,
		"key_audio_map":{"36": 7,"49":8, "51":6}
	},

	{
		"name": "Cherry_G80_3000",
		"display_name": "Cherry G80-3000",
		"files": ["G80-3000.mp3",  "G80-3000_fast1.mp3", "G80-3000_slow1.mp3", "G80-3000_fast2.mp3","G80-3000_slow2.mp3"],
		"non_unique_count": 5,
		"key_audio_map":{"36": 4, "49": 4}
	},

	{
		"name": "Cherry_G80_3494",
		"display_name": "Cherry G80-3494",
		"files": ["G80-3494.mp3", "G80-3494_fast1.mp3", "G80-3494_slow1.mp3", "G80-3494_enter.mp3", "G80-3494_space.mp3", "G80-3494_backspace.mp3"],
		"non_unique_count": 3,
		"key_audio_map":{"36": 3, "49": 4, "51": 5}
	},

	{
		"name": "drum",
		"display_name": "Drum",
		"files": ["1.mp3", "2.mp3", "3.mp3", "4.mp3", "space.mp3", "backspace.mp3", "enter.mp3"],
		"non_unique_count": 4,
		"key_audio_map":{"36": 6, "49": 4, "51": 5}
	},

	{
		"name": "starwars",
		"display_name": "Star Wars",
		"files": ["a.mp3", "b.mp3"],
		"non_unique_count": 1,
		"key_audio_map":{"36": 1}
	}
]`)
}
