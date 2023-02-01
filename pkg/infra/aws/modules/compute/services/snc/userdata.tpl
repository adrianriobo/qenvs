#cloud-config  
rh_subscription:
  username: {{.SubscriptionUsername}}
  password: {{.SubscriptionPassword}}
  auto-attach: true
packages:
  - podman
  - "@virt"
  - jq
  - git
runcmd:
  # Debug libvirt
  #- echo 'log_filters="1:libvirt 1:util 1:qemu"' | tee -a /etc/libvirt/libvirtd.conf
  #- echo 'log_outputs="1:file:/var/log/libvirt/libvirtd.log"' | tee -a /etc/libvirt/libvirtd.conf
  # https://libvirt.org/manpages/libvirtd.html#system-socket-activation
  - echo 'LIBVIRTD_ARGS="--listen"' | tee -a /etc/sysconfig/libvirtd
  - echo 'listen_tls = 0' | tee -a /etc/libvirt/libvirtd.conf
  - echo 'listen_tcp = 1' | tee -a /etc/libvirt/libvirtd.conf
  - echo 'tcp_port = "16509"' | tee -a /etc/libvirt/libvirtd.conf
  - echo 'auth_tcp = "none"' | tee -a /etc/libvirt/libvirtd.conf
  - systemctl daemon-reload 
  - systemctl enable libvirtd-tcp.socket 
  - systemctl start --no-block libvirtd-tcp.socket 
  - systemctl mask libvirtd.socket libvirtd-ro.socket libvirtd-admin.socket libvirtd-tls.socket libvirtd-tcp.socket
  - systemctl enable libvirtd 
  - systemctl start --no-block libvirtd  
  - usermod -a -G libvirt ${username}
  - echo "user.max_user_namespaces=28633" | tee -a /etc/sysctl.d/userns.conf
  - sysctl -p /etc/sysctl.d/userns.conf
  - dnf upgrade -y curl openssl