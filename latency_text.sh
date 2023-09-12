pumba_linux_amd64 netem  --duration 23h --tc-image gaiadocker/iproute2 delay --time 100 kafka.peer0.org1.example.com kafka.peer1.org3.example.com
pumba_linux_amd64 netem  --duration 23h --tc-image gaiadocker/iproute2 delay --time 100 kafka.peer1.org1.example.com kafka.peer1.org3.example.com
pumba_linux_amd64 netem  --duration 23h --tc-image gaiadocker/iproute2 delay --time 100 kafka.peer0.org2.example.com kafka.peer1.org3.example.com
pumba_linux_amd64 netem  --duration 23h --tc-image gaiadocker/iproute2 delay --time 100 kafka.peer1.org2.example.com kafka.peer1.org3.example.com
pumba_linux_amd64 netem  --duration 23h --tc-image gaiadocker/iproute2 delay --time 57 center_kafka