<?php

namespace App\Service\User;

use App\Models\User;
use App\Repositories\Interface\User\UserInterface;
use Illuminate\Database\Eloquent\ModelNotFoundException;

class UserService
{
    public function __construct(
        protected UserInterface $eloquentUserRepository
    ) {}

    public function find(string|int $id, array $with = []): ?User
    {
        $user = $this->eloquentUserRepository->find($id, $with);

        if ( ! $user ) throw new ModelNotFoundException('Usuário não localizado.');

        return $user;
    }

    public function create(array $data): User
    {
        return $this->eloquentUserRepository->create($data);
    }

    public function update(int $id, array $data): User
    {
        return $this->eloquentUserRepository->update($this->find($id), $data);
    }

    public function delete(int $id): void
    {
        $this->eloquentUserRepository->delete($this->find($id));
    }
}
