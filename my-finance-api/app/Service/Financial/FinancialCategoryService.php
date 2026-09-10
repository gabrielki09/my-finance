<?php

namespace App\Service\Financial;

use App\Models\Financial\FinancialCategory;
use App\Repositories\Interface\Financial\FinancialCategoryInterface;
use Illuminate\Database\Eloquent\ModelNotFoundException;

class FinancialCategoryService
{
    public function __construct(
        protected FinancialCategoryInterface $eloquentFinancialCategoryRepository
    ){}

    public function all()
    {
        return $this->eloquentFinancialCategoryRepository->all();
    }

    public function find(string|int $id): FinancialCategory
    {
        $financialCategory = $this->eloquentFinancialCategoryRepository->find($id);

        if ( ! $financialCategory ) throw new ModelNotFoundException('Categoria financeira não localizada.');

        return $financialCategory;
    }

    public function findByUuid(string $uuid): FinancialCategory
    {
        $financialCategory = $this->eloquentFinancialCategoryRepository->findByUuid($uuid);

        if ( ! $financialCategory ) throw new ModelNotFoundException('Categoria financeira não localizada.');

        return $financialCategory;
    }

    public function create(array $data): FinancialCategory
    {
        return $this->eloquentFinancialCategoryRepository->create($data);
    }

    public function update(array $data, string $uuid)
    {
        return $this->eloquentFinancialCategoryRepository->update($this->findByUuid($uuid), $data);
    }

    public function delete(string $uuid)
    {
        $this->eloquentFinancialCategoryRepository->delete($this->findByUuid($uuid));
    }

    public function active(string $uuid)
    {
        $this->eloquentFinancialCategoryRepository->active($this->findByUuid($uuid));
    }
}
