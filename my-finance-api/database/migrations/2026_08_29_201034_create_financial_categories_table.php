<?php

use App\Enum\Financial\FinancialCategoryTypes;
use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create('financial_categories', function (Blueprint $table) {
            $table->id();
            $table->uuid('uuid')->unique('un_financial_category_uuid')->index('idx_financial_category_uuid');
            $table->foreignId('user_id')->constrained('users');
            $table->string('name', 255);
            $table->unique(['name', 'user_id'], 'un_financial_categories_name');
            $table->enum('type', [array_column(FinancialCategoryTypes::cases(), 'value')]);
            $table->string('color', 120)->default('#ffff');
            $table->string('icon', 120)->nullable();
            $table->softDeletes();
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('financial_categories');
    }
};
