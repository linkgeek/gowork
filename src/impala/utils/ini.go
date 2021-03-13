package utils

import "gopkg.in/ini.v1"

type IniParser struct {
	conf_reader *ini.File
}

func (this *IniParser) Load(config_file_name string) error {
	conf, err := ini.Load(config_file_name)
	if err != nil {
		this.conf_reader = nil
		return err
	}
	this.conf_reader = conf
	return nil
}

func (this *IniParser) GetString(section string, key string) string {
	if this.conf_reader == nil {
		return ""
	}

	s := this.conf_reader.Section(section)
	if s == nil {
		return ""
	}

	return s.Key(key).String()
}

func (this *IniParser) GetInt32(section string, key string) int32 {
	if this.conf_reader == nil {
		return 0
	}

	s := this.conf_reader.Section(section)
	if s == nil {
		return 0
	}

	value_int, _ := s.Key(key).Int()

	return int32(value_int)
}