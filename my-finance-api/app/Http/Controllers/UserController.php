<?php

namespace App\Http\Controllers;

use App\Http\Requests\User\CreateUserRequest;
use App\Http\Requests\User\UpdateUserRequest;
use App\Service\User\UserService;

class UserController extends Controller
{
    public function __construct(
        protected UserService $userService
    ){}

    /**
     * Store a newly created resource in storage.
     */
    public function store(CreateUserRequest $request)
    {
        return apiSuccess(
            'Usuário cadastrado com sucesso!',
            [
                'user' => $this->userService->create($request->validated())
            ]
        );
    }

    /**
     * Display the specified resource.
     */
    public function show(string|int $id)
    {
        return apiSuccess(
            'Dados do usuário',
            [
                'user' => $this->userService->find($id)
            ]
        );
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateUserRequest $request, string $id)
    {
        return apiSuccess(
            'Dados do usuário',
            [
                'user' => $this->userService->update($id, $request->validated())
            ]
        );
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        $this->userService->delete($id);
    }
}
