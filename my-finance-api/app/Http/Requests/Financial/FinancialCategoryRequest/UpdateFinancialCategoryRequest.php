<?php

namespace App\Http\Requests\Financial\FinancialCategoryRequest;

use App\Enum\Financial\FinancialCategoryTypes;
use App\Models\Financial\FinancialCategory;
use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Support\Facades\Auth;
use Illuminate\Validation\Rule;

class UpdateFinancialCategoryRequest extends FormRequest
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
        $category = FinancialCategory::query()->where('uuid', $this->route('uuid'))->firstOrFail();

        return [
            'name' => ['sometimes', 'min:3', 'max:255', Rule::unique(FinancialCategory::class, 'name')->ignore($category)],
            'type' => ['sometimes', Rule::enum(FinancialCategoryTypes::class)],
            'color' => ['sometimes', 'max:120'],
            'icon' => ['sometimes', 'max:120'],
        ];
    }

    public function messages()
    {
        return [
            'name.min' => 'O nome da categoria financeira deve conter no minímo :min caracteres.',
            'name.max' => 'O nome da categoria financeira deve conter no máximo :max caracteres.',
            'name.unique' => 'Essa categoria financeira já está cadastrada.',
            'type.enum' => 'O tipo da categoria financeira precisa ser uma Entrada, Saída ou Ambos.',
            'color.max' => 'A cor para exibição da categoria financeira deve conter no máximo :max caracteres.',
            'icon.max' => 'O ícone para exibição da categoria financeira deve conter no máximo :max caracteres.',
        ];
    }
}
