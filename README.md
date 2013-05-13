# Unicode confusables

This Golang library implements the `Skeleton` algorithm from Unicode TR39

See http://www.unicode.org/reports/tr39/

### Examples
```
import "github.com/mtibben/tr39-skeleton-go/unicode/confusables"

confusables.Skeleton("𝔭𝒶ỿ𝕡𝕒ℓ")   # "paypal"

confusables.Confusable("𝔭𝒶ỿ𝕡𝕒ℓ", "paypal")  # true
```
