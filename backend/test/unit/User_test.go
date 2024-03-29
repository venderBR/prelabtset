package unit

import(

	"testing"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"github.com/venderBR/prelabtset/entity"
)

func TestAll(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`test_All` ,func(t *testing.T) {

		user := entity.User{
			Firstname: "phit",
			Lastname: "Ain",
			StudentID: "B6404796",
			Phone: "0645723364",
			Email: "mangpor@gmail.com",
		}

		ok,err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}

func TestFirstname(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`test_required` ,func(t *testing.T) {

		user := entity.User{
			Firstname: "",
			Lastname: "Ain",
			StudentID: "B6404796",
			Phone: "0645723364",
			Email: "mangpor@gmail.com",
		}

		ok,err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Firstname is required"))
	})
}

func TestStudentID(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`test_required` ,func(t *testing.T) {

		user := entity.User{
			Firstname: "phit",
			Lastname: "Ain",
			StudentID: "",
			Phone: "0645723364",
			Email: "mangpor@gmail.com",
		}

		ok,err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("StudentID is required"))
	})
	t.Run(`test_invalid` ,func(t *testing.T) {

		user := entity.User{
			Firstname: "phit",
			Lastname: "Ain",
			StudentID: "n6404796",
			Phone: "0645723364",
			Email: "mangpor@gmail.com",
		}

		ok,err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("StudentID is invalid"))
	})
}

func TestPhone(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`test_required` ,func(t *testing.T) {

		user := entity.User{
			Firstname: "phit",
			Lastname: "Ain",
			StudentID: "b6404796",
			Phone: "",
			Email: "mangpor@gmail.com",
		}

		ok,err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Phone is required"))
	})
	t.Run(`test_invalid` ,func(t *testing.T) {

		user := entity.User{
			Firstname: "phit",
			Lastname: "Ain",
			StudentID: "B6404796",
			Phone: "06457233644",
			Email: "mangpor@gmail.com",
		}

		ok,err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Phone is invalid"))
	})
}