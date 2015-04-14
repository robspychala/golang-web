package shared

import (
    "net"
)

func GetIPAddress() string {

  ipaddress := ""
  addrs, _ := net.InterfaceAddrs()

  for _, a := range addrs {
    if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
      if ipnet.IP.To4() != nil {
        ipaddress = ipnet.IP.String()
      }
    }
  }

  return ipaddress

}
