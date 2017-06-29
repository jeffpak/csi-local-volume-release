package models_test

import (
	. "../.."
	"."

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"time"
	"code.cloudfoundry.org/goshims/osshim/os_fake"
	"code.cloudfoundry.org/goshims/filepathshim/filepath_fake"
	"path"
	"os"
)

var _ = Describe("Controller Service", func() {
	var(
		cs *models.Controller
		context context.Context

		fakeOs       *os_fake.FakeOs
		fakeFilepath *filepath_fake.FakeFilepath
		mountDir     string
		volumeName   string
		volID        *VolumeID
		vc           []*VolumeCapability
		err          error
	)

	BeforeEach(func() {
		mountDir = "/path/to/mount"
		fakeOs = &os_fake.FakeOs{}
		fakeFilepath = &filepath_fake.FakeFilepath{}
		cs = models.NewController(fakeOs, fakeFilepath, mountDir)
		context = &DummyContext{}
		volumeName = "abcd"
		vc = []*VolumeCapability{{Value: &VolumeCapability_Mount{}}}
	})

	Describe("CreateVolume", func() {
		var (
		  response *CreateVolumeResponse
		)

		Context("when CreateVolume is called with a CreateVolumeRequest", func() {
			BeforeEach(func() {
				response = createSuccessful(context, cs, fakeOs, volumeName, vc)
			})

			It("does not fail", func() {
				var (
					mode AccessMode_Mode = AccessMode_UNKNOWN
				  am *AccessMode = &AccessMode{Mode: &mode}
				  volID = &VolumeID{Values: map[string]string{"volume_name": "abcd"}}
				)

				Expect(*response).To(Equal(CreateVolumeResponse{Reply: &CreateVolumeResponse_Result_{
				  Result: &CreateVolumeResponse_Result{
						VolumeInfo: &VolumeInfo{
							AccessMode: am,
							Id: volID},
					}},
				}))
			})
		})

		Context("when a second create is called with the same volume ID", func() {
			BeforeEach(func() {
				createSuccessful(context, cs, fakeOs, volumeName, vc)
			})

			It("does nothing", func() {
				createSuccessful(context, cs, fakeOs, volumeName, vc)
			})
		})

		Context("when create is called without a name", func() {
			BeforeEach(func() {
				volumeName = ""
				_, err = cs.CreateVolume(context, &CreateVolumeRequest{
					Version: &Version{},
					Name: &volumeName,
					VolumeCapabilities: vc,
				})
			})

			It("fails, expecting a name", func() {
				Expect(err.Error()).To(Equal("Missing mandatory 'volume_name'"))
			})
		})
	})

	Describe("DeleteVolume", func() {
		var (
			response *DeleteVolumeResponse
		)

		It("should fail if no volume name is provided", func() {
			_, err := cs.DeleteVolume(context, &DeleteVolumeRequest{
				VolumeId: &VolumeID{},
			})
			Expect(err.Error()).To(Equal("Missing mandatory 'volume_name'"))
		})

		It("should fail if no volume was created", func() {
			volID = &VolumeID{Values: map[string]string{"volume_name": volumeName}}
			_, err := cs.DeleteVolume(context, &DeleteVolumeRequest{
				VolumeId: volID,
			})
			Expect(err.Error()).To(Equal("Volume '" + volumeName + "' not found"))
		})

		Context("when the volume has been created", func() {
			BeforeEach(func() {
				createSuccessful(context, cs, fakeOs, volumeName, vc)
			})

			It("destroys the volume", func() {
				volID = &VolumeID{Values: map[string]string{"volume_name": volumeName}}
				deleteSuccessful(context, cs, *volID)
				Expect(fakeOs.RemoveAllCallCount()).To(Equal(1))

				_, err = cs.DeleteVolume(context, &DeleteVolumeRequest{
					VolumeId: volID,
				})
				Expect(err.Error()).To(Equal("Volume '" + volumeName + "' not found"))
			})

			//Context("when volume has been mounted", func() {
			//	It("/VolumePlugin.Remove unmounts and destroys volume", func() {
			//		mountSuccessful(env, localDriver, volumeId, fakeFilepath)
      //
			//		removeResponse := localDriver.Remove(env, voldriver.RemoveRequest{
			//			Name: volumeId,
			//		})
			//		Expect(removeResponse.Err).To(Equal(""))
			//		Expect(fakeOs.RemoveCallCount()).To(Equal(1))
			//		Expect(fakeOs.RemoveAllCallCount()).To(Equal(1))
      //
			//		getUnsuccessful(env, localDriver, volumeId)
			//	})
			//})
		})

		Context("when the volume has not been created", func() {
			It("returns an error", func() {
				response, err = cs.DeleteVolume(context, &DeleteVolumeRequest{
					VolumeId: volID,
				})
				Expect(err.Error()).To(Equal("Volume '" + volumeName + "' not found"))
			})
		})
	})
})


type DummyContext struct {}

func (*DummyContext) Deadline() (deadline time.Time, ok bool) { return time.Time{}, false }

func (*DummyContext) Done() <-chan struct{} {return nil}

func (*DummyContext) Err() (error){ return nil }

func (*DummyContext) Value(key interface{}) interface{} {return nil}

func createSuccessful(ctx context.Context, cs ControllerServer, fakeOs *os_fake.FakeOs, volumeName string, vc []*VolumeCapability) *CreateVolumeResponse {
	createResponse, err := cs.CreateVolume(ctx, &CreateVolumeRequest{
		Version: &Version{},
		Name: &volumeName,
		VolumeCapabilities: vc,
	})
	Expect(err).To(BeNil())
	Expect(fakeOs.MkdirAllCallCount()).Should(Equal(2))

	volumeDir, fileMode := fakeOs.MkdirAllArgsForCall(1)
	Expect(path.Base(volumeDir)).To(Equal(volumeName))
	Expect(fileMode).To(Equal(os.ModePerm))
	return createResponse
}

func deleteSuccessful(ctx context.Context, cs ControllerServer, volumeID VolumeID) *DeleteVolumeResponse{
	deleteResponse, err := cs.DeleteVolume(ctx, &DeleteVolumeRequest{
		Version: &Version{},
		VolumeId: &volumeID,
	})
	Expect(err).To(BeNil())
	return deleteResponse
}
