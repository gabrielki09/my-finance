<?php

namespace App\Http\Requests\Financial\FinancialReleaseRequest;

use App\Enum\Financial\FinancialReleasesType;
use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Support\Facades\Auth;
use Illuminate\Validation\Rule;

class CreateFinancialReleaseRequest extends FormRequest
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
            'account_id' => ['required', 'exists:financial_accounts,uuid', 'uuid'],
            'category_id' => ['required', 'exists:financial_categories,uuid', 'uuid'],
            'type' => ['required', Rule::enum(FinancialReleasesType::class)],
            'description' => ['required', 'min:3','max:255'],
            'amount' => ['required', 'numeric', 'min:0.01'],
            'due_date' => ['nullable', 'after_or_equal:today'],
            'notes' => ['sometimes', 'string'],
        ];
    }

    public function messages()
    {
        return [
            'account_id.required' => 'A conta financeira é obrigatória.',
            'account_id.exists' => 'A conta financeira precisa ser uma conta válida.',
            'account_id.uuid' => 'O identificador da conta financeira precisa ser um identificador válido.',
            'category_id.required' => 'A categoria financeira é obrigatória.',
            'category_id.exists' => 'A categoria financeira precisa ser uma categoria válida.',
            'category_id.uuid' => 'O identificador da categoria financeira precisa ser um identificador válido.',
            'type.required' => 'O tipo do lançamento financeiro é obrigatório.',
            'type.enum' => 'O tipo do lançamento financeira precisa ser um tipo válido.',
            'description.required' => 'A descrição do lançamento financeiro é obrigatório.',
            'description.min' => 'A descrição do lançamento financeiro deve conter no minímo :min caracteres.',
            'description.max' => 'A descrição do lançamento financeiro deve conter no máximo :max caracteres.',
            'amount.required' => 'O valor do lançamento financeiro é obrigatório.',
            'amount.numeric' => 'O valor do lançamento financeiro deve ser um número válido.',
            'amount.min' => 'O valor do lançamento financeiro deve ser de ao menos R$ :min.',
            'due_date.after_or_equal' => 'A data de vencimento quando definida, não pode ser anterior ao dia atual.',
            'notes.string' => 'As notas do lançamento financeiro devem ser um texto válido.',
        ];
    }
}
