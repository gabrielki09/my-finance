<?php

namespace App\Enum\Financial;

enum FinancialReleasesStatus: string
{
    case PENDING = 'pending';
    case PAID = 'paid';
    case CANCELLED = 'cancelled';
    case OVERDUE = 'overdue';
}
