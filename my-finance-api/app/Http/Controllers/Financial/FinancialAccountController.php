<?php

namespace App\Http\Controllers\Financial;

use App\Http\Controllers\Controller;
use App\Http\Requests\Financial\FinancialCategoryRequest\CreateFinancialAccountRequest;
use App\Http\Requests\Financial\FinancialCategoryRequest\UpdateFinancialAccountRequest;
use App\Service\Financial\FinancialAccountService;

class FinancialAccountController extends Controller
{
    public function __construct(
        protected FinancialAccountService $financialAccountService
    ){}
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        return apiSuccess(
            'Todas as contas financeiras',
            $this->financialAccountService->all()
        );
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(CreateFinancialAccountRequest $request)
    {
        return apiSuccess(
            'Conta financeira cadastrada com sucesso!',
            $this->financialAccountService->create($request->validated())
        );
    }

    /**
     * Display the specified resource.
     */
    public function show(string $uuid)
    {
        return apiSuccess(
            'Dados da conta financeira',
            $this->financialAccountService->findByUuid($uuid)
        );
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateFinancialAccountRequest $request, string $uuid)
    {
        return apiSuccess(
            'Conta financeira alterada com sucesso!',
            $this->financialAccountService->update($request->validated(), $uuid)
        );
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $uuid)
    {
        $this->financialAccountService->delete($uuid);
    }
}
