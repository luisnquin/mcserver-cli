package app

func (v *Version) GetServer(name string) (Pod, error) {
	s, ok := v.Servers[name]
	if !ok {
		return nil, ErrServerNotFound
	}

	s.config = v.config
	s.version = v.name
	s.name = name
	s.saver = v

	return s, nil
}

func (v *Version) NewServer(name string) error {
	if _, ok := v.Servers[name]; ok {
		return ErrServerAlreadyExists
	}

	v.Servers[name] = new(Server)

	return v.saveData()
}

func (v *Version) CopyServer(target, name string) error {
	if _, ok := v.Servers[name]; ok {
		return ErrServerAlreadyExists
	}

	s, ok := v.Servers[target]
	if !ok {
		return ErrServerNotFound
	}

	s.IsCopy = true
	v.Servers[name] = s

	return v.saveData()
}

func (v *Version) DeleteServer(name string) error {
	if _, ok := v.Servers[name]; !ok {
		return ErrServerNotFound
	}

	delete(v.Servers, name)

	return v.saveData()
}
