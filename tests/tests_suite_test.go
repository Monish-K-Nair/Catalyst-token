package invite_token_test_test

// import (
// 	"testing"

// 	. "github.com/onsi/ginkgo/v2"
// 	. "github.com/onsi/gomega"
// 	utils "catalyst-token/utils"

// )

// func TestTests(t *testing.T) {
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "Tests Suite")
// }

// var _ = Describe("Invite tokens", func() {
// 	Context("initially", func() {

// 		// RetrieveAPI
// 		It("has 0 items", func() {
// 			Expect((utils.APIClient())).Should(BeZero())
// 		})
// 		// Random token Validate must fail
// 		It("has 0 items", func() {
// 			Expect((utils.APIClient())).Should(BeZero())
// 		})
// 	})

// 	Context("when a new token is added", func() {
// 		// Retrieve API and compute count
// 		oldCount := utils.APIClient()
// 		// generae token without admin credentails must fail

// 		// login with wrong credentilas must fail

// 		// Generate invite token API with admin credentials
// 		// tk.AddItem(tk1)

// 		Context("the invite tokens", func() {
// 			It("has 1 more unique token than it had earlier", func() {
// 					// Retrieve API and compute count
// 					newCount := utils.APIClient()
// 				Expect(newCount).Should(Equal(oldCount + 1))
// 			})
// 		})
// 	})

// 	Context("that has 1 Token", func() {

// 		// Generate invite token API
// 		// tk.AddItem(tk1)
// 				// Generate invite token API
// 		// tk.AddItem(tk1)
// 				// Generate invite token API
// 		// tk.AddItem(tk1)


// 		oldCount := utils.APIClient()

// 		Context("removing 1 unit item A", func() {
// 			// Generate invite token API
// 			// tk.AddItem(tk1)

// 			It("should reduce the number of items by 1", func() {
// 				newCount := utils.APIClient()
// 				Expect(newCount).Should(Equal(oldCount - 1))
// 			})
// 		})
// 	})
// 	Context("that has 1 Token", func() {
// 		tk.AddItem(tk1)

// 		oldCount := utils.APIClient()

// 		Context("removing 1 unit item A", func() {
// 			tk.RemoveItem(tk1.Token, 1)

// 			It("should reduce the number of items by 1", func() {
// 				Expect(tk.TotalUniqueItems()).Should(Equal(oldCount - 1))
// 			})
// 		})
// 	})

// })