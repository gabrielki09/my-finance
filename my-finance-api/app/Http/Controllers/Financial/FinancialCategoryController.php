<?php

namespace App\Http\Controllers\Financial;

use App\Http\Controllers\Controller;
use App\Http\Requests\Financial\FinancialCategoryRequest\CreateFinancialCategorRequest;
use App\Service\Financial\FinancialCategoryService;
use Illuminate\Http\Request;

class FinancialCategoryController extends Controller
{
    public function __construct(
        protected FinancialCategoryService $financialCategoryService
    ){}
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        return apiSuccess(
            'Todas as categorias financeiras',
            $this->financialCategoryService->all()
        );
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(CreateFinancialCategorRequest $request)
    {
        return apiSuccess(
            'Categoria financeira cadastrada com sucesso!',
            $this->financialCategoryService->create($request->validated())
        );
    }

    /**
     * Display the specified resource.
     */
    public function show(string $uuid)
    {
        return apiSuccess(
            'Dados da categoria financeira.',
            $this->financialCategoryService->findByUuid($uuid)
        );
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, string $uuid)
    {
        return apiSuccess(
            'Categoria financeira alterada com sucesso!',
            $this->financialCategoryService->update($request->validated(), $uuid)
        );
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $uuid)
    {
        $this->financialCategoryService->delete($uuid);
    }
}
