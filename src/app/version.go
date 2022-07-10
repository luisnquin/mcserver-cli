package app

func (v *Version) GetServer(name string) (Server, error) {
	s, ok := v.Servers[name]
	if !ok {
		return s, ErrServerNotFound
	}

	s.config = v.config
	s.version = v.name
	s.name = name

	return s, nil
}

func (v *Version) NewServer(name string) error {
	if _, ok := v.Servers[name]; ok {
		return ErrServerAlreadyExists
	}

	v.Servers[name] = Server{}

	return nil
}

func (v *Version) CopyServer(sTarget, name string) error {
	if _, ok := v.Servers[name]; ok {
		return ErrServerAlreadyExists
	}

	s, ok := v.Servers[sTarget]
	if !ok {
		return ErrServerNotFound
	}

	s.IsCopy = true

	v.Servers[name] = s

	return nil
}

func (v *Version) DeleteServer(name string) error {
	if _, ok := v.Servers[name]; !ok {
		return ErrServerNotFound
	}

	delete(v.Servers, name)

	return nil
}