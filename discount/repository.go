package discount

import (
	"context"

	"github.com/uptrace/bun"
)

type Repository interface {
	GetDiscount(ctx context.Context, productName string) (*Coupon, error)
	CreateDiscount(ctx context.Context, dto CreateCouponDTO) (*Coupon, error)
	UpdateDiscount(ctx context.Context, c Coupon) error
	DeleteDiscount(ctx context.Context, productName string) error
}

type PGRepository struct {
	db *bun.DB
}

func NewPGRepository(db *bun.DB) Repository {
	return &PGRepository{db: db}
}

func (p *PGRepository) GetDiscount(ctx context.Context, productName string) (*Coupon, error) {
	c := CouponEntity{ProductName: productName}
	err := p.db.NewSelect().Model(&c).Where("product_name = ?", productName).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &Coupon{
		ID:          c.ID,
		ProductName: c.ProductName,
		Description: c.Description,
		Amount:      c.Amount,
	}, nil
}

func (p *PGRepository) CreateDiscount(ctx context.Context, dto CreateCouponDTO) (*Coupon, error) {
	c := CouponEntity{
		ID:          0,
		ProductName: dto.ProductName,
		Description: dto.Description,
		Amount:      dto.Amount,
	}
	var id int
	_, err := p.db.NewInsert().Model(&c).Returning("id").Exec(ctx, &id)
	if err != nil {
		return nil, err
	}
	return &Coupon{
		ID:          id,
		ProductName: dto.ProductName,
		Description: dto.Description,
		Amount:      dto.Amount,
	}, nil
}

func (p *PGRepository) UpdateDiscount(ctx context.Context, c Coupon) error {
	e := CouponEntity{
		ID:          c.ID,
		ProductName: c.ProductName,
		Description: c.Description,
		Amount:      c.Amount,
	}
	_, err := p.db.NewUpdate().Model(&e).WherePK().ExcludeColumn().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (p *PGRepository) DeleteDiscount(ctx context.Context, productName string) error {
	_, err := p.db.NewDelete().Model(&CouponEntity{}).Where("product_name = ?", productName).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
