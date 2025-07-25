---
output: sys-report-{{ raw "date --iso-8601" }}.md
---

# 💥 System Information Report for {{ sh "hostname" }}

* **Hostname**: {{ sh "hostname" }}
* **Username**: {{ sh "whoami" }}
* **Uptime**: {{ sh "uptime -p" }}
* **System**: {{ sh "uname -a" }}
* **CPU**: {{ sh "uname -m" }} — {{ sh "nproc" }} cores
* **IP Address**: {{ sh "hostname -I || ip a | grep inet" }}
* **Default Gateway**: {{ sh "ip route | grep default || netstat -rn | grep default" }}

---

## 🧠 Memory Usage

{{ shell "free -h" }}

---

## 📉 Disk Usage

{{ shell "df -h /" }}

---

Generated on: {{ sh "date" }}
