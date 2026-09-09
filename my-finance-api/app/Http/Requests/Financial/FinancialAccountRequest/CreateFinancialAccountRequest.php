<?php

namespace App\Http\Requests\Financial\FinancialCategoryRequest;

use App\Models\Financial\FinancialCategory;
use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Support\Facades\Auth;
use Illuminate\Validation\Rule;

class CreateFinancialAccountRequest extends FormRequest
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
            'user_id' => ['required', 'exists:users,id'],
            'name' => ['required', 'min:3', 'max:255', Rule::unique(FinancialCategory::class, 'name')]
        ];
    }

    public function messages()
    {
        return [
            'user_id.required' => 'O identificador do usuário responsável é obrigatório.',
            'user_id.exists' => 'O usuário responsável deve ser um usuário válido.',
            'name.required' => 'O nome da categoria financeira é obrigatório.',
            'name.min' => 'O nome da categoria financeira deve conter no minímo :min caracteres.',
            'name.max' => 'O nome da categoria financeira deve conter no máximo :max caracteres.',
            'name.unique' => 'Essa categoria financeira já está cadastrada.',
        ];
    }
}
