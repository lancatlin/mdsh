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
