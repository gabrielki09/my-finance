<?php

namespace App\Enum\Financial;

enum FinancialCategoryTypes: string
{
    case INCOME = 'income';
    case EXPENSE = 'expense';
    case BOTH = 'both';
}
