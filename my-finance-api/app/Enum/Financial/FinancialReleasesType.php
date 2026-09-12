<?php

namespace App\Enum\Financial;

enum FinancialReleasesType: string
{
    case INCOME = 'income';
    case EXPENSE = 'expense';
    case TRANSFER = 'transfer';
}
