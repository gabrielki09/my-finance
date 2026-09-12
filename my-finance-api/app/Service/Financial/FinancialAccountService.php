<?php

namespace App\Service\Financial;

use App\Models\Financial\FinancialAccount;
use App\Repositories\Interface\Financial\FinancialAccountInterface;
use Illuminate\Database\Eloquent\ModelNotFoundException;

class FinancialAccountService
{
    public function __construct(
        protected FinancialAccountInterface $eloquentFinancialAccountRepository
    ){}

    public function all()
    {
        return $this->eloquentFinancialAccountRepository->all();
    }

        public function find(string|int $id, int $userId): FinancialAccount
    {
        $financialAccount = $this->eloquentFinancialAccountRepository->find($id, $userId);

        if ( ! $financialAccount ) throw new ModelNotFoundException('Categoria financeira não localizada.');

        return $financialAccount;
    }

    public function findByUuid(string $uuid, int $userId): FinancialAccount
    {
        $financialAccount = $this->eloquentFinancialAccountRepository->findByUuid($uuid, $userId);

        if ( ! $financialAccount ) throw new ModelNotFoundException('Conta financeira não localizada.');

        return $financialAccount;
    }

    public function create(array $data): FinancialAccount
    {
        return $this->eloquentFinancialAccountRepository->create([
            'user_id' => $data['user_id'],
            'name' => $data['name'],
            'type' => $data['type'],
            'initial_balance' => $data['initial_balance'],
            'current_balance' => $data['initial_balance'],
        ]);
    }

    public function update(array $data, string $uuid)
    {
        return $this->eloquentFinancialAccountRepository->update($this->findByUuid($uuid, $data['user_id']), $data);
    }

    public function delete(string $uuid, int $userId)
    {
        $this->eloquentFinancialAccountRepository->delete($this->findByUuid($uuid, $userId));
    }

    public function active(string $uuid, int $userId)
    {
        $this->eloquentFinancialAccountRepository->active($this->findByUuid($uuid, $userId));
    }
}
