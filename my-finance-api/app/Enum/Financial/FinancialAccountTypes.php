<?php

namespace App\Enum\Financial;

enum FinancialAccountTypes: string
{
    case CASH = 'cash';
    case CHECKING_ACCOUNT = 'checking_account';
    case SAVINGS_ACCOUNT = 'savings_account';
    case DIGITAL_ACCOUNT = 'digital_account';
    case CREDIT_CARD = 'credit_card';
    case INVESTMENT = 'investment';
    case OTHER = 'other';
}
