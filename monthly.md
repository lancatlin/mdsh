---
output: "reports_{{.from}}.md"
params:
  from:
    required: true
  to:
    required: true
  file:
    default: ~/.hledger.journal
---

# Finance Report from {{.from}} to {{.to}}

## Previous status:

{{ sh "hledger -f $file balancesheet -e $from" }}

## Expenses:

### Expense Summary by Category

{{ sh "hledger -f $file bal expenses -b $from -e $to --depth 3 --sort amount" }}

{{ sh "hledger -f $file bal expenses -b $from -e $to --depth 2 --sort amount" }}

## Net Income:

{{sh "hledger -f $file income -b $from -e $to --monthly"}}

## Ending balance:

{{sh "hledger -f $file balancesheet -e $to"}}
