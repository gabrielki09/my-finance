<?php

namespace App\Http\Controllers\Financial;

use App\Http\Controllers\Controller;
use App\Http\Requests\Financial\FinancialReleaseRequest\CreateFinancialReleaseRequest;
use App\Http\Requests\Financial\FinancialReleaseRequest\UpdateFinancialReleaseRequest;
use App\Service\Financial\FinancialReleaseService;
use Illuminate\Http\Request;

class FinancialReleaseController extends Controller
{
    public function __construct(
        protected FinancialReleaseService $financialReleaseService
    ){}

    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        return apiSuccess(
            'Todas os lançamentos financeiros',
            $this->financialReleaseService->all()
        );
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(CreateFinancialReleaseRequest $request)
    {
        return apiSuccess(
            'Lançamento financeiro cadastrado com sucesso!',
            $this->financialReleaseService->create(buildArrayDataWithUserId($request)),
            201
        );
    }

    /**
     * Display the specified resource.
     */
    public function show(Request $req, string $uuid)
    {
        return apiSuccess(
            'Dados do lançamento financeiro',
            $this->financialReleaseService->findByUuid($uuid, $req->user()->id)
        );
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(UpdateFinancialReleaseRequest $request, string $uuid)
    {
        return apiSuccess(
            'Lançamento financeiro atualizado com sucesso!',
            $this->financialReleaseService->update(buildArrayDataWithUserId($request), $uuid)
        );
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(Request $req, string $uuid)
    {
        $this->financialReleaseService->delete($uuid, getUserIdByRequest($req));
    }
}
