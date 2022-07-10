package app

func (v *Version) GetServer(name string) (Server, error) {
	s, ok := v.Servers[name]
	if !ok {
		return s, ErrServerNotFound
	}

	return s, nil
}

func (v *Version) AddServer(name string) error {
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
