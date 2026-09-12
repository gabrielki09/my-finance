<?php

namespace App\Http\Controllers\Financial;

use App\Http\Controllers\Controller;
use App\Http\Requests\Financial\FinancialCategoryRequest\CreateFinancialCategoryRequest;
use App\Http\Requests\Financial\FinancialCategoryRequest\UpdateFinancialCategoryRequest;
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
    public function store(CreateFinancialCategoryRequest $request)
    {
        return apiSuccess(
            'Categoria financeira cadastrada com sucesso!',
            $this->financialCategoryService->create(buildArrayDataWithUserId($request)),
            201
        );
    }

    /**
     * Display the specified resource.
     */
    public function show(Request $req, string $uuid)
    {
        return apiSuccess(
            'Dados da categoria financeira.',
            $this->financialCategoryService->findByUuid($uuid, getUserIdByRequest($req))
        );
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateFinancialCategoryRequest $request, string $uuid)
    {
        return apiSuccess(
            'Categoria financeira alterada com sucesso!',
            $this->financialCategoryService->update(buildArrayDataWithUserId($request), $uuid)
        );
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Request $req, string $uuid)
    {
        $this->financialCategoryService->delete($uuid, getUserIdByRequest($req));
    }

    public function active(Request $req, string $uuid)
    {
        $this->financialCategoryService->active($uuid, getUserIdByRequest($req));
    }
}
