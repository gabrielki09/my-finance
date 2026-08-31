<?php

namespace App\Repositories\Eloquent;

use App\Repositories\Interface\BaseInterface;
use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Collection;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Pagination\LengthAwarePaginator;

abstract class EloquentBaseRepository implements BaseInterface
{
    protected Model $model;

    public function __construct()
    {
        $this->model = app($this->modelClass());
    }

    abstract protected function modelClass(): string;

    public function query(): Builder
    {
        return $this->model->newQuery();
    }

    public function all(array $columns = ['*'], array $with = []): Collection
    {
        return $this->query()->with($with)->get($columns);
    }

    public function paginate(
        int $perPage = 15,
        array $columns = ['*'],
        array $with = []
    ): LengthAwarePaginator {
        return $this->query()
                ->with($with)
                ->paginate($perPage, $columns);
    }

    public function find(string|int $id, array $with = []): ?Model
    {
        return $this->query()
                    ->with($with)
                    ->find($id);
    }

    public function findByUuid(string $uuid, array $with = []): ?Model
    {
        return $this->query()
                    ->with($with)
                    ->where('uuid', $uuid)
                    ->first();
    }

    public function create(array $data): Model
    {
        return $this->query()->create($data);
    }

    public function update(Model $model, array $data): Model
    {
        $model->fill($data);
        $model->save();

        return $model->refresh();
    }

    public function delete(Model $model): void
    {
        $model->delete();
    }

    public function active(Model $model): void
    {
        $model->restore();
    }

}
