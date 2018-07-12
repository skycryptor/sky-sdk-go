//
// Created by Tigran on 7/4/18.
//

#ifndef CRYPTOMAIC_CAPSULE_H
#define CRYPTOMAIC_CAPSULE_H

#include "Point.h"

namespace SkyCryptor {

  /**
   * \brief Combination of parameters as a definition for cryptographic capsule
   * Each capsule contains E(Point), V(Point), s(BigNumber)
   */
  class Capsule {
    /// Defining Capsule particles
    Point particleE;
    Point particleV;
    BigNumber particleS;
    Point particleXG;

    /// Keeping crypto context available for capsule
    /// NOTE: this class is not taking responsibility for cleaning up this pointer
    Context *context;

    bool reEncrypted = false;
   public:
    /**
     * \brief Making capsule with given particles
     * @param E
     * @param V
     * @param S
     * @param ctx
     */
    Capsule(Point& E, Point& V, BigNumber& S, Context *ctx);
    /**
     * \brief Making capsule with particles and public key to be encoded with it
     * @param E
     * @param V
     * @param S
     * @param XG
     * @param ctx
     */
    Capsule(Point& E, Point& V, BigNumber& S, Point& XG, Context *ctx, bool isReEncription = false);
    /**
     * \brief Copy constructor from another capsule
     * @param other
     */
    Capsule(const Capsule& other);
    ~Capsule() = default;

    /**
     * Getting particle E as a Point
     * @return
     */
    Point get_particleE() const;

    /**
     * Getting particle V as a Point
     * @return
     */
    Point get_particleV() const;

    /**
     * Getting particle S as a BigNumber
     * @return
     */
    BigNumber get_particleS() const;

    /**
     * Getting particle XG
     * @return
     */
    Point get_particleXG() const;

    /**
     * \brief Setting capsule as re-encryption capsule
     */
    void setReEncrypted();

    /**
     * \brief Checking if we have re-encryption capsule or not
     * @return
     */
    bool isreEncrypted();

    /**
     * \brief Serializing capsule to bytes
     * @return
     */
    vector<char> toBytes();

    /**
     * \brief Getting Capsule from encoded bytes
     * @param buffer
     * @param length
     * @param ctx
     * @return
     */
    static Capsule from_bytes(const char *buffer, int length, Context *ctx);
    static Capsule from_bytes(vector<char> buffer, Context *ctx);
  };
}

#endif //CRYPTOMAIC_CAPSULE_H
