<?php

namespace App\Service\Auth;

use App\Models\User;
use App\Repositories\Interface\User\UserInterface;
use Exception;
use Illuminate\Database\Eloquent\ModelNotFoundException;
use Illuminate\Support\Facades\Hash;

class AuthService
{
    public function __construct(
        protected UserInterface $eloquentUserRepository
    ){}

    private function generateToken(User $user): string
    {
        return $user->createToken('api')->plainTextToken;
    }

    private function checkPassword(string $informedPassword, User $user): bool
    {
        return Hash::check($informedPassword, $user->password);
    }

    private function findByEmail(string $email): ?User
    {
        $user = $this->eloquentUserRepository->findByEmail($email);

        if ( ! $user ) throw new ModelNotFoundException('Credencias incorretas.');

        return $user;
    }

    public function login(string $email, string $password): array
    {
        $user = $this->findByEmail($email);
        if ( !$this->checkPassword($password, $user) ) throw new Exception('Credencias incorretas.');

        return [
            'token' => $this->generateToken($user)
        ];
    }
}
