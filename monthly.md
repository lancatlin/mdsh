---
name: "reports/${from}.md"
params:
  from:
    required: true
  to:
    required: true
  file:
    default: ~/.hledger.journal
---

# Finance Report from {from} to {to}

## Previous status:

{{hledger -f $file balancesheet -e $from }}

## Expenses:

### Expense Summary by Category

{{hledger -f $file bal expenses -b $from -e $to --depth 3 --sort amount}}

{{hledger -f $file bal expenses -b $from -e $to --depth 2 --sort amount}}

## Net Income:

{{hledger -f $file income -b $from -e $to --monthly}}

## Ending balance:

{{hledger -f $file balancesheet -e $to}}
