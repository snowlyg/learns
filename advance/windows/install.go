package windows

import (
	"fmt"

	"golang.org/x/sys/windows/svc/eventlog"
	"golang.org/x/sys/windows/svc/mgr"
)

func ServiceInstall(svcName, execPath, dispalyName string, args ...string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(svcName)
	if err == nil {
		s.Close()
		return fmt.Errorf("service %s already exists", svcName)
	}
	config := mgr.Config{
		DisplayName: dispalyName,
		StartType:   mgr.StartAutomatic,
	}
	if len(args) >= 2 {
		config.ServiceStartName = args[0]
		config.Password = args[1]
	}
	s, err = m.CreateService(svcName, execPath, config)
	if err != nil {
		return fmt.Errorf("CreateService() failed: %s", err)
	}
	defer s.Close()
	err = eventlog.InstallAsEventCreate(svcName, eventlog.Error|eventlog.Warning|eventlog.Info)
	if err != nil {
		s.Delete()
		return fmt.Errorf("InstallAsEventCreate() failed: %s", err)
	}
	return nil
}
