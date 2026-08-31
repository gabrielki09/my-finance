<?php

namespace App\Repositories\Interface\User;

use App\Models\User;
use Illuminate\Database\Eloquent\Builder;

interface UserInterface
{
    public function query(): Builder;
    public function find(string|int $id, array $with = []): ?User;
    public function findByEmail(string $email): ?User;
    public function create(array $data): User;
    public function update(User $user, array $data): User;
    public function delete(User $user): void;
}
