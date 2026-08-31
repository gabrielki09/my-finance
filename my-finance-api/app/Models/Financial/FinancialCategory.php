<?php

namespace App\Models\Financial;

use Illuminate\Database\Eloquent\Attributes\Fillable;
use Illuminate\Database\Eloquent\Concerns\HasUuids;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

#[Fillable([
    'uuid',
    'user_id',
    'name',
    'type',
    'color',
    'icon'
])]
class FinancialCategory extends Model
{
    use HasUuids, SoftDeletes;

    protected $primaryKey = 'id';

    protected function casts()
    {
        return [
            'deleted_at' => 'date'
        ];
    }
}
