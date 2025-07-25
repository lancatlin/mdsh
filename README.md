# mdsh: Run Shell Scripts in Markdown Templates

Need repetitive document generation from CLI outputs? mdsh is what you need.

Define your template in markdown format, and generate them on the fly!

## Installation

requirements: `go` installed

```sh
git clone github.com/lancatlin/mdsh && cd mdsh
go build .
./mdsh examples/system-info.md
```

It will execute this template:

~~~~markdown
---
output: "sys-report.md"
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
~~~~

Then you will get `sys-report.md` like:

~~~~markdown
# 💥 System Information Report for `fedora`

* **Hostname**: `fedora`
* **Username**: `wancat`
* **Uptime**: `up 1 hour, 59 minutes`
* **System**: `Linux fedora 6.15.6-200.fc42.x86_64 #1 SMP PREEMPT_DYNAMIC Thu Jul 10 15:22:32 UTC 2025 x86_64 GNU/Linux`
* **CPU**: `x86_64` — `16` cores
* **IP Address**: `172.26.198.115 2405:dc00:ec83:ec80:af9c:87ed:9bae:bd0d`
* **Default Gateway**: `default via 172.26.198.50 dev wlp1s0 proto dhcp src 172.26.198.115 metric 600`

---

## 🧠 Memory Usage


```
               total        used        free      shared  buff/cache   available
Mem:            30Gi       7.5Gi        17Gi       234Mi       6.5Gi        23Gi
Swap:          8.0Gi          0B       8.0Gi

```


---

## 📉 Disk Usage


```
Filesystem                                             Size  Used Avail Use% Mounted on
/dev/mapper/luks-d53d6eca-6a17-4120-bb6c-05c501334c6d   93G   58G   34G  64% /

```


---

Generated on: `Fri 25 Jul 2025 12:42:26 AEST`
~~~~

---

## Custom Parameters from command line

You can define custom parameters in the frontmatter, and pass the value through command arguments.

~~~~markdown
---
output: "monthly_report_{{.b}}.md"
params:
  b:
    required: true
    default: 2011-01
    usage: |
      Required.
      The begin time of report.
      Examples:
        2025
        2025-07
        2025-07-15
        Jul
  e:
    required: true
    default: 2011-02
    usage: |
      Required.
      The end time of report.
      Same format as -b
  f:
    default: examples/ledger.j
    usage: The ledger file to parse
---

# Finance Report from {{.b}} to {{.e}}

## Previous status:

{{ shell "hledger -f $f balancesheet -e $b" }}

## Expenses:

### Expense Summary by Category

{{ shell "hledger -f $f bal expenses -b $b -e $e --depth 3 --sort amount --budget" }}

## Net Income:

{{shell "hledger -f $f income -b $b -e $e --monthly"}}

## Ending balance:

{{shell "hledger -f $f balancesheet -e $e"}}
~~~~

Run the template:

**Note: you need to install `hledger` to run the example**

```
mdsh examples/monthly.md -b 2011-01 -e 2011-02
```

The result:

~~~~markdown
# Finance Report from 2011-01 to 2011-02

## Previous status:

`Balance Sheet 2010-12-31

                                || 2010-12-31
================================++============
 Assets                         ||
--------------------------------++------------
 Assets:Checking                ||   $-225.00
--------------------------------++------------
                                ||   $-225.00
================================++============
 Liabilities                    ||
--------------------------------++------------
 Liabilities:Mortgage:Principal ||   $-200.00
--------------------------------++------------
                                ||   $-200.00
================================++============
 Net:                           ||    $-25.00`

## Expenses:

### Expense Summary by Category

`Budget performance in 2011-01:

               ||                          Jan
===============++==============================
 Expenses      || $5,629.00
 Expenses:Food ||   $109.00 [ 109% of $100.00]
---------------++------------------------------
               || $5,629.00 [5629% of $100.00]`

`$5,500.00  Expenses:Auto
             $109.00  Expenses:Food
              $20.00  Expenses:Books
--------------------
           $5,629.00`

## Net Income:

`Income Statement 2011-01

                         ||        Jan
=========================++============
 Revenues                ||
-------------------------++------------
 Income:Salary           ||  $2,000.00
-------------------------++------------
                         ||  $2,000.00
=========================++============
 Expenses                ||
-------------------------++------------
 Expenses:Auto           ||  $5,500.00
 Expenses:Books          ||     $20.00
 Expenses:Food:Groceries ||    $109.00
-------------------------++------------
                         ||  $5,629.00
=========================++============
 Net:                    || $-3,629.00`

## Ending balance:

`Balance Sheet 2011-01-31

                                || 2011-01-31
================================++============
 Assets                         ||
--------------------------------++------------
 Assets:Checking                ||  $1,366.00
 Assets:Savings                 || $-5,200.00
--------------------------------++------------
                                || $-3,834.00
================================++============
 Liabilities                    ||
--------------------------------++------------
 Liabilities:MasterCard         ||     $20.00
 Liabilities:Mortgage:Principal ||   $-200.00
--------------------------------++------------
                                ||   $-180.00
================================++============
 Net:                           || $-3,654.00`
~~~~
