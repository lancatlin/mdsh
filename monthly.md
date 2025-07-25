---
output: "reports_{{.b}}.md"
params:
  b:
    required: true
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
    usage: |
      Required.
      The end time of report.
      Same format as -b
  f:
    default: ~/.hledger.journal
    usage: The ledger file to parse
---

# Finance Report from {{.b}} to {{.e}}

## Previous status:

{{ sh "hledger -f $f balancesheet -e $b" }}

## Expenses:

### Expense Summary by Category

{{ sh "hledger -f $f bal expenses -b $b -e $e --depth 3 --sort amount" }}

{{ sh "hledger -f $f bal expenses -b $b -e $e --depth 2 --sort amount" }}

## Net Income:

{{sh "hledger -f $f income -b $b -e $e --monthly"}}

## Ending balance:

{{sh "hledger -f $f balancesheet -e $e"}}
