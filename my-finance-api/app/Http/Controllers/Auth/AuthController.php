<?php

namespace App\Http\Controllers\Auth;

use App\Http\Controllers\Controller;
use App\Http\Requests\Auth\AuthRequest;
use App\Service\Auth\AuthService;
use Illuminate\Http\Request;

class AuthController extends Controller
{
    public function __construct(
        protected AuthService $authService
    ){}

    public function login(AuthRequest $req)
    {
        $data = $req->validated();

        return apiSuccess(
            'Login bem sucedido!',
            $this->authService->login($data['email'], $data['password'])
        );
    }

    public function logout(Request $req)
    {
        $req->user()->currentAccessToken()->delete();

        return apiSuccess(
            'Logout bem sucedido!'
        );
    }

    public function me(Request $req)
    {
        return apiSuccess(
            'Meus dados',
            [
                'user' => $req->user()
            ]
        );
    }
}
