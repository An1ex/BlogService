package setting

var sections = make(map[string]interface{})

func (vs *ViperSetting) ReadSection(k string, v interface{}) error {
	err := vs.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (vs *ViperSetting) ReloadAllSection() error {
	for k, v := range sections {
		err := vs.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
