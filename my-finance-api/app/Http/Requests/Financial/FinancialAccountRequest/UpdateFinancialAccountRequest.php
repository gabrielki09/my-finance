<?php

namespace App\Http\Requests\Financial\FinancialCategoryRequest;

use App\Enum\Financial\FinancialAccountTypes;
use App\Models\Financial\FinancialCategory;
use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Support\Facades\Auth;
use Illuminate\Validation\Rule;

class UpdateFinancialAccountRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return Auth::check();
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            'name' => ['sometimes', 'min:3', 'max:255', Rule::unique(FinancialCategory::class, 'name')->ignore($this->route('uuid'))],
            'type' => ['sometimes', Rule::enum(FinancialAccountTypes::class)],
            'initial_balance' => ['sometimes', 'numeric', 'min:0.01']
        ];
    }

    public function messages()
    {
        return [
            'name.min' => 'O nome da categoria financeira deve conter no minímo :min caracteres.',
            'name.max' => 'O nome da categoria financeira deve conter no máximo :max caracteres.',
            'name.unique' => 'Essa categoria financeira já está cadastrada.',
            'type.enum' => 'O tipo da conta fincaneira precisa ser uma conta válida.',
            'initial_balance.numeric' => 'O saldo inicial da conta financeira deve ser um número válido.',
            'initial_balance.min' => 'O saldo inicial da conta financeira deve ser pelo R$ :min .',
        ];
    }
}
