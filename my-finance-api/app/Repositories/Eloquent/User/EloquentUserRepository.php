<?php

namespace App\Repositories\Eloquent\User;

use App\Models\User;
use App\Repositories\Interface\User\UserInterface;
use Illuminate\Database\Eloquent\Builder;

class EloquentUserRepository implements UserInterface
{
    public function query(): Builder
    {
        return User::query();
    }

    public function find(string|int $id, array $with = []): ?User
    {
        return $this->query()->with($with)->find($id);
    }

    public function findByEmail(string $email): ?User
    {
        return $this->query()->where('email', $email)->first();
    }

    public function create(array $data): User
    {
        return $this->query()->create($data);
    }

    public function update(User $user, array $data): User
    {
        $user->fill($data);
        $user->save();

        return $user->refresh();
    }

    public function delete(User $user): void
    {
        $user->delete();
    }
}
