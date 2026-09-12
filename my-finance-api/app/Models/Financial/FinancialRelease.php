<?php

namespace App\Models\Financial;

use App\Enum\Financial\FinancialReleasesStatus;
use App\Enum\Financial\FinancialReleasesType;
use Illuminate\Database\Eloquent\Attributes\Fillable;
use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

#[Fillable([
    'uuid',
    'user_id',
    'account_id',
    'category_id',
    'type',
    'description',
    'amount',
    'transaction_date',
    'due_date',
    'paid_date',
    'status',
    'notes',
])]
class FinancialRelease extends Model
{
    use HasUuids, SoftDeletes;

    protected $primaryKey = 'uuid';

    public function uniqueIds()
    {
        return ['uuid'];
    }

    protected function casts()
    {
        return [
            'status' => FinancialReleasesStatus::class,
            'type' => FinancialReleasesType::class,
            'amount' => 'decimal:2',
            'deleted_at' => 'date'
        ];
    }
}
