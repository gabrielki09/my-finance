<?php

namespace App\Models\Financial;

use App\Enum\Financial\FinancialAccountTypes;
use Illuminate\Database\Eloquent\Attributes\Fillable;
use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

#[Fillable([
    'uuid',
    'user_id',
    'name',
    'type',
    'initial_balance',
    'current_balance',
])]
class FinancialAccount extends Model
{
    use HasUuids, SoftDeletes;

    protected $primaryKey = 'id';

    protected function casts()
    {
        return [
            'type' => FinancialAccountTypes::class,
            'initial_balance' => 'decimal:2',
            'current_balance' => 'decimal:2',
            'deleted_at' => 'date',
        ];
    }
}
