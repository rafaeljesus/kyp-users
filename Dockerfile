FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>
ADD kyp-users /kyp-users
ENV KYP_USERS_DB="postgres://postgres:@docker/kyp_users_dev?sslmode=disable"
ENV KYP_USERS_PORT="3000"
ENV KYP_SECRET_KEY="c91267c27a8599ca0480ea505487d052e3b63a1dd39819db853225a518200399"
ENTRYPOINT ["/kyp-users"]
