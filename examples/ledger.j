; -*- ledger -*-
; This example file retrived from ledger-cli.org example journal file
; https://ledger-cli.org/doc/ledger3.html#Example-Journal-File

= /^Income/
  (Liabilities:Tithe)                    0.12

~ Monthly
  Assets:Checking                     $500.00
  Income:Salary

~ Monthly
   Expenses:Food  $100
   Assets

2010/12/01 * Checking balance
  Assets:Checking                   $1,000.00
  Equity:Opening Balances

2010/12/20 * Organic Co-op
  Expenses:Food:Groceries             $ 37.50  ; [=2011/01/01]
  Expenses:Food:Groceries             $ 37.50  ; [=2011/02/01]
  Expenses:Food:Groceries             $ 37.50  ; [=2011/03/01]
  Expenses:Food:Groceries             $ 37.50  ; [=2011/04/01]
  Expenses:Food:Groceries             $ 37.50  ; [=2011/05/01]
  Expenses:Food:Groceries             $ 37.50  ; [=2011/06/01]
  Assets:Checking                   $ -225.00

2010/12/28=2011/01/01 Acme Mortgage
  Liabilities:Mortgage:Principal    $  200.00
  Expenses:Interest:Mortgage        $  500.00
  Expenses:Escrow                   $  300.00
  Assets:Checking                  $ -1000.00

2011/01/02 Grocery Store
  Expenses:Food:Groceries             $ 65.00
  Assets:Checking

2011/01/05 Employer
  Assets:Checking                   $ 2000.00
  Income:Salary

2011/01/14 Bank
  ; Regular monthly savings transfer
  Assets:Savings                     $ 300.00
  Assets:Checking

2011/01/19 Grocery Store
  Expenses:Food:Groceries             $ 44.00  ; hastag: not block
  Assets:Checking

2011/01/25 Bank
  ; Transfer to cover car purchase
  Assets:Checking                  $ 5,500.00
  Assets:Savings
  ; :nobudget:

apply tag hastag: true
apply tag nestedtag: true
2011/01/25 Tom's Used Cars
  Expenses:Auto                    $ 5,500.00
  ; :nobudget:
  Assets:Checking

2011/01/27 Book Store
  Expenses:Books                       $20.00
  Liabilities:MasterCard
end tag
2011/12/01 Sale
  Assets:Checking:Business            $ 30.00
  Income:Sales
end tag
