delete FROM registry where key in ('order_disallow_user_cancel',"domain_file_server_prefix");

/** 2023-02-17 09:50 */
ALTER TABLE IF EXISTS public.pay_order
    RENAME final_fee TO final_amount;

ALTER TABLE IF EXISTS public.pay_order
    RENAME paid_fee TO paid_amount;