<?php

namespace App\Repositories\Interface;

use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Collection;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Pagination\LengthAwarePaginator;

interface BaseInterface
{
    public function query(): Builder;
    public function all(array $columns = ['*'], array $with = []): Collection;
    public function paginate(
        int $perPage = 15,
        array $columns = ['*'],
        array $with = []
    ): LengthAwarePaginator;

    public function find(string|int $id, int $userId, array $with = []): ?Model;
    public function findByUuid(string $uuid, int $userId, array $with = []): ?Model;
    public function create(array $data): Model;
    public function update(Model $model, array $data): Model;
    public function delete(Model $model): void;
    public function active(Model $model): void;
}
